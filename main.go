package main

import (
	"fmt"
	"os"
	"p2p-chat/internal/chat"
)

// main function - kept small on purpose

func main() {

	args := os.Args

	if len(args) > 1 {
		switch args[1] {
		case "listen":
			chat.Listen(args[2])
		case "connect":
			chat.Connect(args[2])
		default:
			fmt.Println("Usage: go run . [listen|connect] address")
		}
	} else {
		fmt.Println("Usage: go run . [listen|connect] address")
	}
}
