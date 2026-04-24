package main

import (
	"fmt"
	"os"
	"p2p-chat/internal/chat"
)

func main() {

	args := os.Args

	if len(args) > 1 {
		if args[1] == "listen" {
			chat.Listen(args[2])
		} else if args[1] == "connect" {
			chat.Connect(args[2])
		} else {
			fmt.Println("Usage: go run . [listen|connect] address")
		}
	} else {
		fmt.Println("Usage: go run . [listen|connect] address")
	}
}
