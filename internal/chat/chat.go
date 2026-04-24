package chat

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"p2p-chat/crypto"
	"p2p-chat/internal/transport"
)

func Listen(address string) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on " + address)
	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")
	handle(conn)
}

func Connect(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to " + address)
	handle(conn)
}

func handle(conn net.Conn) {
	sharedKey, err := crypto.ExchangeKeys(conn, conn)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			encrypted, err := transport.ReceiveFrame(conn)
			if err != nil {
				if err == io.EOF {
					fmt.Println("Disconnected")
				} else {
					fmt.Println("Read error:", err)
				}
			}
			decrypted, err := crypto.Decrypt(encrypted, sharedKey)
			if err != nil {
				if err == io.EOF {
					fmt.Println("Disconnected")
				} else {
					fmt.Println("Warning:", err)
					continue
				}
			}
			fmt.Println("friend>", string(decrypted))
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		ciphertext, err := crypto.Encrypt([]byte(text), sharedKey)
		if err != nil {
			panic(err)
		}
		err = transport.SendFrame(conn, ciphertext)
		if err != nil {
			panic(err)
		}
	}
}
