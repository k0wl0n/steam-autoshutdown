# Steam Auto Shutdown

<img width="1600" height="621" alt="image" src="https://github.com/user-attachments/assets/76a72abd-ea30-4ffe-80ee-2b4aac832ec7" />


This is a robust command-line tool written in Go that automatically shuts down your computer when Steam (or any other application) finishes downloading large files.

It monitors your global network traffic. When it detects a sustained period of high speed (downloading) followed by a period of low speed (finished), it triggers a system shutdown.

## Supported Platforms

- **macOS** (AppleScript / `shutdown`)
- **Linux** (`shutdown`)
- **Windows** (`shutdown.exe`)

## How It Works

1.  **Waiting**: The tool sits quietly and watches your network speed.
2.  **Downloading**: When the speed exceeds the start threshold (default 500 KB/s), it knows a download has begun.
3.  **Finished**: When the speed drops below the stop threshold (default 100 KB/s) and stays there for the idle duration (default 60 seconds), it assumes the download is complete.
4.  **Shutdown**: It first attempts a polite shutdown (asking apps to quit). If that fails, it schedules a forced shutdown.

## Installation

### Homebrew (Recommended)

```bash
brew tap k0wl0n/tap
brew install steam-autoshutdown
```

### Build from Source
1.  Open Terminal in this directory.
2.  Build the binary:
    ```bash
    go build -o steamshutdown
    ```

## Usage

Run the tool with `sudo` (required for shutdown permissions):

```bash
sudo ./steamshutdown
```

### Customization (Flags)

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--start-threshold` | `-s` | `500` | Network speed (KB/s) to identify a download has started. |
| `--stop-threshold` | `-e` | `100` | Network speed (KB/s) to identify a download has finished. |
| `--idle-duration` | `-d` | `60` | Duration (seconds) of low traffic before shutting down. |
| `--interface` | `-i` | `""` | Specific network interface (e.g., `en0`). Default: All non-loopback. |
| `--dry-run` | | `false` | Simulate the process without actually shutting down. |

### Demo Output

Here is an example of a dry-run execution:

```text
$ sudo ./steamshutdown --dry-run --start-threshold 10 --stop-threshold 5 --idle-duration 5
Starting Steam Auto Shutdown Monitor...
-------------------------------------
Configuration:
  Start Threshold: 10 KB/s
  Stop Threshold:  5 KB/s
  Idle Duration:   5 seconds
  Interface:       ALL (excluding loopback)
  Mode:            DRY RUN (No actual shutdown)
-------------------------------------
Waiting for download to start... Current speed: 12015.21 KB/s   
Download detected! Speed: 12015.21 KB/s. Monitoring for completion...
Downloading... Speed: 12003.15 KB/s
```

Here is an example of a live run where the download dips but recovers before finally finishing:

```text
$ sudo ./steamshutdown --start-threshold 10
Password:
Starting Steam Auto Shutdown Monitor...
-------------------------------------
Configuration:
  Start Threshold: 10 KB/s
  Stop Threshold:  100 KB/s
  Idle Duration:   60 seconds
  Interface:       ALL (excluding loopback)
-------------------------------------
Waiting for download to start... Current speed: 11998.78 KB/s
Download detected! Speed: 11998.78 KB/s. Monitoring for completion...
Speed low (29.20 KB/s). Shutting down in 57 seconds...
Speed recovered (552.26 KB/s). Resuming download watch.
Speed low (2.61 KB/s). Shutting down in 45 seconds...
```

**Example:**
Wait for a download to start (threshold 1MB/s), and shut down only if speed stays below 50KB/s for 2 minutes:

```bash
sudo ./steamshutdown --start-threshold 1024 --stop-threshold 50 --idle-duration 120
```

## Troubleshooting

-   **Permission Denied**: Ensure you run the command with `sudo`.
-   **Immediate Shutdown**: If you are not downloading anything when you start the tool, it will stay in the "Waiting" state. It only shuts down *after* it has seen a download start and then finish.
