# Usage

Run the tool with `sudo` to ensure it has permission to shut down the system.

```bash
sudo ./steamshutdown [flags]
```

## Common Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--start-threshold` | `-s` | `500` | Network speed (KB/s) to identify a download has started. |
| `--stop-threshold` | `-e` | `100` | Network speed (KB/s) to identify a download has finished. |
| `--idle-duration` | `-d` | `60` | Duration (seconds) of low traffic before shutting down. |
| `--interface` | `-i` | `""` | Specific network interface (e.g., `en0`). If empty, monitors all non-loopback interfaces. |
| `--dry-run` | | `false` | Simulate the process without actually shutting down. |

## Demo Output

Here is what the tool looks like in action (running in dry-run mode):

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

## Examples

### Basic Usage
Monitor all interfaces, wait for >500KB/s to start, shutdown if <100KB/s for 60s.
```bash
sudo ./steamshutdown
```

### Specific Interface (Wi-Fi)
Only monitor `en0` (common for macOS Wi-Fi).
```bash
sudo ./steamshutdown --interface en0
```

### Safe Test (Dry Run)
Test with lower thresholds to verify logic.
```bash
sudo ./steamshutdown --dry-run --start-threshold 10 --stop-threshold 5 --idle-duration 10
```
