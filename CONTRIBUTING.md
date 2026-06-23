# Contributing to p2p-chat

Thanks for your interest in contributing! Here's how to get started.

## Setup

1. **Go 1.26.2** — If you don't have it, run `sh install-go.sh` or download from [go.dev](https://go.dev/dl/).
2. Clone the repo and `cd` into it.
3. Run `go mod download` to fetch dependencies.

## Development

```sh
# Format and vet before running
sh run.sh

# Or manually:
go fmt ./...
go vet ./...
go run . listen localhost:5555
```

## Pull Requests

- Keep changes small and focused.
- Run `go fmt ./...` and `go vet ./...` before committing.
- Open your PR against the `master` branch.
- Describe what the change does and why it's needed.

## Reporting Issues

Found a bug or have a feature idea? Open an issue with:

- What you expected to happen
- What actually happened
- Steps to reproduce (if applicable)

## Code Style

- Follow standard Go conventions (`gofmt`, `go vet`).
- Keep the codebase simple — this is a learning project, so readability matters more than cleverness.

Thanks for contributing!
