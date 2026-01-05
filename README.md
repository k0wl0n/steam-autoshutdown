# Steam Auto Shutdown (Golang Edition)

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

### Prerequisites
- Go 1.20 or higher (if building from source)

### Build
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

You can customize the thresholds using flags:

-   `--start-threshold`: Speed in KB/s to trigger "Downloading" state.
-   `--stop-threshold`: Speed in KB/s to consider "Idle".
-   `--idle-duration`: Seconds to wait in idle state before shutting down.

**Example:**
Wait for a download to start (threshold 1MB/s), and shut down only if speed stays below 50KB/s for 2 minutes:

```bash
sudo ./steamshutdown --start-threshold 1024 --stop-threshold 50 --idle-duration 120
```

## Troubleshooting

-   **Permission Denied**: Ensure you run the command with `sudo`.
-   **Immediate Shutdown**: If you are not downloading anything when you start the tool, it will stay in the "Waiting" state. It only shuts down *after* it has seen a download start and then finish.
