# Installation

## Prerequisites

- Go 1.24 or higher (if building from source)
- `sudo` privileges (required for system shutdown)

## Building from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/k0wl0n/steam-mac-autoshutdown.git
   cd steam-mac-autoshutdown
   ```

2. Build the binary using Taskfile (recommended):
   ```bash
   task build
   ```
   Or manually with Go:
   ```bash
   go build -o dist/steamshutdown main.go
   ```

## Binary Download

Check the [Releases](https://github.com/k0wl0n/steam-mac-autoshutdown/releases) page for pre-built binaries for your operating system.
