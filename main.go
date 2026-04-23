package main

import "fmt"
import "os"

func main() {

	args := os.Args

	if (len(args) > 1) {
		if (args[1] == "listen") {
			fmt.Println("Listening...")
		} else if (args[1] == "connect") {
			fmt.Println("Connecting...")
		} else {
			fmt.Println("Usage: go run . [listen|connect] address")
		}
	} else {
		fmt.Println("Usage: go run . [listen|connect] address")
	}
}
