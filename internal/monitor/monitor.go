package monitor

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/k0wl0n/steam-autoshutdown/internal/system"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

// Config holds the settings for the monitor.
type Config struct {
	DownloadThresholdKB int
	IdleThresholdKB     int
	IdleDurationSeconds int
	DryRun              bool
	InterfaceName       string
	CheckSteam          bool
}

// State represents the current status of the download monitor.
type State string

const (
	StateWaiting     State = "WAITING"
	StateDownloading State = "DOWNLOADING"
)

// Start begins the network monitoring process.
// It blocks until the shutdown condition is met or an error occurs.
func Start(cfg Config) error {
	fmt.Println("-------------------------------------")
	fmt.Printf("Configuration:\n")
	fmt.Printf("  Start Threshold: %d KB/s\n", cfg.DownloadThresholdKB)
	fmt.Printf("  Stop Threshold:  %d KB/s\n", cfg.IdleThresholdKB)
	fmt.Printf("  Idle Duration:   %d seconds\n", cfg.IdleDurationSeconds)
	if cfg.InterfaceName != "" {
		fmt.Printf("  Interface:       %s\n", cfg.InterfaceName)
	} else {
		fmt.Println("  Interface:       ALL (excluding loopback)")
	}
	if cfg.DryRun {
		fmt.Println("  Mode:            DRY RUN (No actual shutdown)")
	}
	if cfg.CheckSteam {
		fmt.Println("  Steam Check:     ENABLED")
	}
	fmt.Println("-------------------------------------")

	currentState := StateWaiting
	var idleStartTime time.Time
	isIdle := false

	// Helper to calculate total bytes based on config
	getTotalBytes := func() (uint64, error) {
		// We use true to get per-interface stats, so we can filter
		counters, err := net.IOCounters(true)
		if err != nil {
			return 0, err
		}

		var total uint64
		for _, stat := range counters {
			// If specific interface requested, only count that
			if cfg.InterfaceName != "" {
				if stat.Name == cfg.InterfaceName {
					return stat.BytesRecv, nil
				}
				continue
			}

			// Otherwise sum all NON-loopback interfaces
			// loopback usually named "lo0", "lo", "localhost"
			if strings.HasPrefix(strings.ToLower(stat.Name), "lo") {
				continue
			}
			total += stat.BytesRecv
		}

		// If specific interface was requested but not found, return 0 (or handle error)
		return total, nil
	}

	// Get initial network counters
	lastTotalBytesRecv, err := getTotalBytes()
	if err != nil {
		return fmt.Errorf("failed to get network stats: %w", err)
	}
	lastTime := time.Now()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Helper to check if Steam is running
	isSteamRunning := func() bool {
		processes, err := process.Processes()
		if err != nil {
			return false
		}
		for _, p := range processes {
			name, err := p.Name()
			if err != nil {
				continue
			}
			// Check for "steam" (linux), "steam.exe" (windows), or "steam_osx" (macOS)
			if strings.EqualFold(name, "steam") || strings.EqualFold(name, "steam.exe") || strings.EqualFold(name, "steam_osx") {
				return true
			}
		}
		return false
	}

	for range ticker.C {
		// Optional: Check if Steam is running if configured
		if cfg.CheckSteam {
			if !isSteamRunning() {
				fmt.Printf("\r[Info] Steam process not found. Waiting...                 ")
				// We don't exit, just wait until Steam starts or user stops
				// Alternatively, we could exit if the user wants strictly Steam-only monitoring
				// But standard behavior is usually "monitor traffic", check process as a gate

				// Reset state if Steam closes mid-download
				if currentState != StateWaiting {
					currentState = StateWaiting
					fmt.Printf("\nSteam closed. Resetting to waiting state.\n")
				}
				continue
			}
		}

		// Get current network counters
		totalBytesRecv, err := getTotalBytes()
		if err != nil {
			fmt.Printf("Warning: Failed to read network stats: %v\n", err)
			continue
		}

		currentTime := time.Now()
		duration := currentTime.Sub(lastTime).Seconds()

		// Calculate speed in KB/s
		// Handle potential wrap around or reset if needed (though unlikely with uint64 in short time)
		var bytesDiff float64
		if totalBytesRecv >= lastTotalBytesRecv {
			bytesDiff = float64(totalBytesRecv - lastTotalBytesRecv)
		} else {
			// Interface stats might have reset
			bytesDiff = 0
		}

		speedKB := (bytesDiff / 1024) / duration

		// Update last stats for next iteration
		lastTotalBytesRecv = totalBytesRecv
		lastTime = currentTime

		// State Machine Logic
		switch currentState {
		case StateWaiting:
			fmt.Printf("\rWaiting for download to start... Current speed: %.2f KB/s   ", speedKB)
			if speedKB > float64(cfg.DownloadThresholdKB) {
				currentState = StateDownloading
				fmt.Printf("\nDownload detected! Speed: %.2f KB/s. Monitoring for completion...\n", speedKB)
			}

		case StateDownloading:
			if speedKB < float64(cfg.IdleThresholdKB) {
				if !isIdle {
					isIdle = true
					idleStartTime = time.Now()
				}

				elapsedIdle := time.Since(idleStartTime).Seconds()
				remaining := float64(cfg.IdleDurationSeconds) - elapsedIdle

				fmt.Printf("\rSpeed low (%.2f KB/s). Shutting down in %.0f seconds...   ", speedKB, math.Max(0, remaining))

				if elapsedIdle >= float64(cfg.IdleDurationSeconds) {
					fmt.Printf("\nDownload finished (speed < %d KB/s for %d seconds).\n", cfg.IdleThresholdKB, cfg.IdleDurationSeconds)
					return system.Shutdown(cfg.DryRun)
				}
			} else {
				// Traffic picked up again
				if isIdle {
					fmt.Printf("\nSpeed recovered (%.2f KB/s). Resuming download watch.\n", speedKB)
					isIdle = false
				}
				fmt.Printf("\rDownloading... Speed: %.2f KB/s            ", speedKB)
			}
		}
	}

	return nil
}
