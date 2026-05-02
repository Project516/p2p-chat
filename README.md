# p2p chat

Simple peer to peer chat written in go.

## Requirements

* go 1.26.2

## Usage

Run `go run . listen localhost:5555` to start a server. Then run `go run . connect localhost:5555` to connect to it!

### Commands

* `/nick`: change display name (for yourself on other users screen)
* `/quit`: quit the chat
* `/help`: display available commands
