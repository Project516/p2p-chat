package main

import (
	"fmt"
	"os"
	"p2p-chat/internal/chat"
)

// main function - kept small on purpose

func main() {

	var version = "1.0.0-alpha"

	args := os.Args

	// pass values from cli
	if len(args) > 1 {
		switch args[1] {
		case "listen":
			// fallback if not port given
			if len(args) == 2 {
				fmt.Println("Error: No port given")
				fmt.Println("Defaulting to port :5555")
				chat.Listen("localhost:5555")
			}
			chat.Listen(args[2])
		case "connect":
			chat.Connect(args[2])
		case "version":
			fmt.Println("p2p-chat version: " + version)
		default:
			fmt.Println("Usage: go run . [listen|connect] address")
		}
	} else {
		fmt.Println("Usage: go run . [listen|connect] address")
	}
}
