# Installation

## Prerequisites

- Go 1.24 or higher (if building from source)
- `sudo` privileges (required for system shutdown)

## Homebrew (macOS/Linux)

The easiest way to install is via Homebrew:

```bash
brew tap k0wl0n/tap
brew install steam-autoshutdown
```

## Building from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/k0wl0n/steam-autoshutdown.git
   cd steam-autoshutdown
   ```

2. Build the binary using Taskfile (recommended):
   ```bash
   task build
   ```
   Or manually with Go:
   ```bash
   go build -o steamshutdown main.go
   ```

## Binary Download

Check the [Releases](https://github.com/k0wl0n/steam-autoshutdown/releases) page for pre-built binaries for your operating system.
