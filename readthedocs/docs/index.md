# Steam Auto Shutdown

**Steam Auto Shutdown** is a CLI tool for macOS (and Linux/Windows) that monitors network traffic and shuts down the computer when downloads finish.

It is designed to replicate the functionality of [Steam Auto Shutdown](https://github.com/diogomartino/steam-auto-shutdown) but with a focus on macOS support and command-line flexibility.

## Features

- **Network Monitoring**: Detects when downloads start and stop based on traffic speed.
- **Smart Detection**: Only shuts down after a sustained period of idleness to avoid false positives.
- **Cross-Platform**: Built with Go, works on macOS, Linux, and Windows.
- **Dry Run Mode**: Test your configuration safely without actually shutting down.
- **Specific Interface**: Monitor a specific network interface (e.g., Wi-Fi only).
