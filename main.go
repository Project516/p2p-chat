package main

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"golang.org/x/crypto/nacl/box"
	"io"
	"net"
	"os"
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
	pub, priv, err := box.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	_, err = conn.Write(pub[:])
	if err != nil {
		panic(err)
	}
	var peerPub [32]byte
	_, err = io.ReadFull(conn, peerPub[:])
	if err != nil {
		panic(err)
	}
	var sharedKey [32]byte
	box.Precompute(&sharedKey, &peerPub, priv)
	fmt.Println("Encryption established.")
	go func() {
		for {
			var lenBuf [2]byte
			_, err := io.ReadFull(conn, lenBuf[:])
			if err != nil {
				if err == io.EOF {
					fmt.Println("Disconnected")
				} else {
					fmt.Println("Read error: ", err)
				}
				os.Exit(0)
			}
			length := binary.BigEndian.Uint16(lenBuf[:])
			encrypted := make([]byte, length)
			_, err = io.ReadFull(conn, encrypted)
			if err != nil {
				if err == io.EOF {
					fmt.Println("Disconnected")
				} else {
					fmt.Println("Read error: ", err)
				}
				os.Exit(0)
			}
			var nonce [24]byte
			copy(nonce[:], encrypted[:24])
			ciphertext := encrypted[24:]
			decrypted, ok := box.OpenAfterPrecomputation(nil, ciphertext, &nonce, &sharedKey)
			if !ok {
				fmt.Println("Warning: failed to decrypt message!")
				continue
			}
			fmt.Println("friend>", string(decrypted))
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		var nonce [24]byte
		_, err := io.ReadFull(rand.Reader, nonce[:])
		if err != nil {
			panic(err)
		}
		ciphertext := box.SealAfterPrecomputation(nonce[:], []byte(text), &nonce, &sharedKey)
		length := uint16(len(ciphertext))
		lenBuf := make([]byte, 2)
		binary.BigEndian.PutUint16(lenBuf, length)
		_, err = conn.Write(lenBuf)
		if err != nil {
			panic(err)
		}
		_, err = conn.Write(ciphertext)
		if err != nil {
			panic(err)
		}
	}
}
