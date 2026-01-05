# Usage

Run the tool with `sudo` to ensure it has permission to shut down the system.

```bash
sudo ./dist/steamshutdown [flags]
```

## Common Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--start-threshold` | `-s` | `500` | Speed (KB/s) to trigger "Downloading" state. |
| `--stop-threshold` | `-e` | `100` | Speed (KB/s) to consider "Finished/Idle". |
| `--idle-duration` | `-d` | `60` | Seconds to wait in idle state before shutting down. |
| `--interface` | `-i` | `""` | Specific network interface (e.g., `en0`). Default: All non-loopback. |
| `--dry-run` | | `false` | Simulate process without shutting down. |

## Examples

### Basic Usage
Monitor all interfaces, wait for >500KB/s to start, shutdown if <100KB/s for 60s.
```bash
sudo ./dist/steamshutdown
```

### Specific Interface (Wi-Fi)
Only monitor `en0` (common for macOS Wi-Fi).
```bash
sudo ./dist/steamshutdown --interface en0
```

### Safe Test (Dry Run)
Test with lower thresholds to verify logic.
```bash
sudo ./dist/steamshutdown --dry-run --start-threshold 10 --stop-threshold 5 --idle-duration 10
```
