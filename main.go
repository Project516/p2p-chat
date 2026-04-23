package main

import "fmt"
import "os"
import "net"

func main() {

	args := os.Args

	if (len(args) > 1) {
		if (args[1] == "listen") {
			listen(args[2])
		} else if (args[1] == "connect") {
			fmt.Println("Connecting...")
		} else {
			fmt.Println("Usage: go run . [listen|connect] address")
		}
	} else {
		fmt.Println("Usage: go run . [listen|connect] address")
	}
}

func listen(address string) {
	net.Listen("tcp", address)
	fmt.Println("Listening on " + address)
}