//go:build linux

package system

import (
	"fmt"
	"os/exec"
)

// Shutdown initiates a system shutdown on Linux.
// It uses the standard 'shutdown' command.
func Shutdown(dryRun bool) error {
	fmt.Println("Initiating shutdown sequence (Linux)...")

	if dryRun {
		fmt.Println("[DRY RUN] Shutdown command would be executed here.")
		fmt.Println("[DRY RUN] System would execute 'sudo shutdown -h +1'.")
		return nil
	}

	// Schedule shutdown in 1 minute
	// Most Linux distros support "+1" for 1 minute delay
	cmd := exec.Command("sudo", "shutdown", "-h", "+1")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute shutdown command: %w", err)
	}

	fmt.Println("Shutdown scheduled in 1 minute. Run 'sudo shutdown -c' to cancel.")
	return nil
}
