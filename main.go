package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"p2p-chat/crypto"
	"p2p-chat/internal/transport"
)

func main() {

	args := os.Args

	if len(args) > 1 {
		if args[1] == "listen" {
			listen(args[2])
		} else if args[1] == "connect" {
			connect(args[2])
		} else {
			fmt.Println("Usage: go run . [listen|connect] address")
		}
	} else {
		fmt.Println("Usage: go run . [listen|connect] address")
	}
}

func listen(address string) {
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

func connect(address string) {
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
			os.Exit(0)
			decrypted, err := crypto.Decrypt(encrypted, sharedKey)
			if err != nil {
				if err == io.EOF {
					fmt.Println("Disconnected")
				} else {
					fmt.Println("Warning:", err)
					continue
				}
			}
			os.Exit(0)
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
