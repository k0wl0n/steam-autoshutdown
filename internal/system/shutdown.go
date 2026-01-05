package system

import (
	"fmt"
	"os/exec"
)

// Shutdown initiates a system shutdown on macOS.
// It first attempts a graceful shutdown using AppleScript (System Events).
// If that fails, it falls back to the 'shutdown' command which requires root privileges.
func Shutdown(dryRun bool) error {
	fmt.Println("Initiating shutdown sequence...")

	if dryRun {
		fmt.Println("[DRY RUN] Shutdown command would be executed here.")
		fmt.Println("[DRY RUN] System would attempt AppleScript shutdown, then 'sudo shutdown -h +1'.")
		return nil
	}

	// Attempt 1: AppleScript (Graceful)
	// This tells the Finder/System to shut down, allowing apps to save state.
	script := `tell application "System Events" to shut down`
	cmd := exec.Command("osascript", "-e", script)
	
	if err := cmd.Run(); err == nil {
		fmt.Println("Shutdown command sent via AppleScript.")
		return nil
	} else {
		fmt.Println("AppleScript shutdown failed. Attempting forced shutdown...")
	}

	// Attempt 2: System Command (Forceful)
	// "sudo shutdown -h +1" schedules a shutdown in 1 minute.
	// We use "+1" to give the user a brief moment to cancel if they are watching.
	cmd = exec.Command("sudo", "shutdown", "-h", "+1")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute shutdown command: %w", err)
	}

	fmt.Println("Shutdown scheduled in 1 minute. Run 'sudo killall shutdown' to cancel.")
	return nil
}
