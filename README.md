# p2p-chat

Simple peer-to-peer encrypted chat written in Go.

Connections are established directly between two peers using TCP, with all messages encrypted using NaCl's `box` primitive (Curve25519 + XSalsa20 + Poly1305) via a key exchange at connection time. No server or intermediary is required — just two instances that connect to each other.

## Features

- **End-to-end encryption** — Uses NaCl's authenticated encryption (`box.SealAfterPrecomputation`) with a Curve25519 key exchange
- **Peer-to-peer** — Direct TCP connection between exactly two peers; no relay server needed
- **Framed messaging** — Length-prefixed binary framing for reliable message delivery over TCP
- **Nicknames** — Change your display name at any time with `/nick`
- **Version command** — Check the current build version with `/version`

## Requirements

- Go 1.26.2 or later

## Usage

**Start a listener** (waits for an incoming connection):

```bash
go run . listen localhost:5555
```

**Connect to a listener** (from another terminal or machine):

```bash
go run . connect localhost:5555
```

The listener and connector roles are symmetric once connected — both peers can send and receive messages.

### Quick-start scripts

| Script       | Command              |
|--------------|----------------------|
| `host.sh`    | Listens on `localhost:5555` |
| `connect.sh` | Connects to `localhost:5555` |
| `run.sh`     | Formats, vets, and runs the project |

> **Note:** `run.sh` runs `go fmt`, `go vet`, and then `go run .` without arguments, which will print usage instructions. Pass `listen` or `connect` subcommands explicitly.

### Commands (in chat)

- `/nick <name>` — Change your display name
- `/quit` — Leave the chat
- `/version` — Display program version
- `/help` — Show available commands

## How It Works

1. **Key exchange** — On connection, both peers generate a Curve25519 keypair, exchange public keys, and compute a shared secret using `box.Precompute`.
2. **Encryption** — Every message is encrypted with the shared key using a random 24-byte nonce and XSalsa20-Poly1305.
3. **Framing** — Messages are sent over TCP with a 2-byte big-endian length prefix, so the receiver knows exactly how many bytes to read.

## Project Structure

```
.
├── main.go                    # Entry point and CLI argument handling
├── crypto/
│   └── crypto.go              # Key exchange, encrypt, decrypt (NaCl box)
├── internal/
│   ├── chat/
│   │   └── chat.go            # Listen, Connect, and chat loop with commands
│   ├── transport/
│   │   └── transport.go       # Length-prefixed framing for TCP messages
│   └── version/
│       └── version.go         # Reads version from assets/version.txt
├── assets/
│   └── version.txt            # Current version string
├── host.sh                    # Start a listener
├── connect.sh                 # Connect to a listener
└── run.sh                     # Format, vet, and run
```

## Security Considerations

- The key exchange does **not** authenticate the peer — it provides encryption but not identity verification. An active man-in-the-middle could intercept the key exchange.
- The chat supports exactly **two peers** per session.
- Nonces are generated using `crypto/rand` and are unique per message.
- The shared key is ephemeral and only lives for the duration of the connection.

## License

This project is licensed under the GNU Affero General Public License v3.0. See [LICENSE](LICENSE) for details.
