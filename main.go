package main

import "fmt"
import "os"
import "net"
import "bufio"

func main() {

	args := os.Args

	if (len(args) > 1) {
		if (args[1] == "listen") {
			listen(args[2])
		} else if (args[1] == "connect") {
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
	conn, err := ln.Accept();
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
	fmt.Println("Connected")
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			text := scanner.Text()
			fmt.Println("friend>", text)
		}
		fmt.Println("Disconnected")
		os.Exit(0)
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Fprintln(conn, text)
	}
}