package chat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"p2p-chat/crypto"
	"p2p-chat/internal/transport"
	"strings"
)

// Listen function

func Listen(address string) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on " + address)
	conn, err := ln.Accept()
	if err != nil {
		log.Printf("accept error: %v", err)
		return
	}
	fmt.Println("Connected")
	handle(conn)
}

// Connect function

func Connect(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to " + address)
	handle(conn)
}

// Handle function - during chat

func handle(conn net.Conn) {
	defer conn.Close()
	sharedKey, err := crypto.ExchangeKeys(conn, conn)
	if err != nil {
		panic(err)
	}
	var nick string = "friend"
	prompt := "> "
	go func() {
		for {
			encrypted, err := transport.ReceiveFrame(conn)
			if err != nil {
				if err == io.EOF {
					fmt.Println("Disconnected")
				} else {
					fmt.Println("Read error:", err)
				}
				os.Exit(0)
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
			fmt.Printf("\n%s\n>", string(decrypted))
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "/nick") { // nickname command
			if len(text) < 7 {
				fmt.Println("Usage: /nick <nickname>")
			} else {
				newNick := strings.TrimSpace(text[6:])
				if newNick == "" {
					fmt.Println("Nickname cannot be empty.")
				} else {
					nickChangeAlert := "! " + nick + " changed their nickname to " + newNick
					ciphertext, err := crypto.Encrypt([]byte(nickChangeAlert), sharedKey)
					if err != nil {
						panic(err)
					}
					err = transport.SendFrame(conn, ciphertext)
					if err != nil {
						panic(err)
					}
					nick = newNick
					fmt.Print("> ")
					continue
				}
			}
		} else if strings.HasPrefix(text, "/quit") { // quit command
			ciphertext, err := crypto.Encrypt([]byte("!"+nick+" left the chat"), sharedKey)
			if err != nil {
				panic(err)
			}
			err = transport.SendFrame(conn, ciphertext)
			if err != nil {
				panic(err)
			}
			fmt.Println("Disconnecting...")
			os.Exit(0)
		} else if strings.HasPrefix(text, "/help") { // help command
			fmt.Print("\n! help menu:\n! run /nick to change nickname\n! run /quit to leave chat\n! run /help to display this menu\n\n>")
			continue
		}
		messageText := text
		messageText = nick + "> " + text
		ciphertext, err := crypto.Encrypt([]byte(messageText), sharedKey)
		if err != nil {
			panic(err)
		}
		err = transport.SendFrame(conn, ciphertext)
		if err != nil {
			panic(err)
		}
		fmt.Print(prompt)
	}
}
