//go:build windows

package system

import (
	"fmt"
	"os/exec"
)

// Shutdown initiates a system shutdown on Windows.
// It uses the 'shutdown' command with a 60-second delay.
func Shutdown(dryRun bool) error {
	fmt.Println("Initiating shutdown sequence (Windows)...")

	if dryRun {
		fmt.Println("[DRY RUN] Shutdown command would be executed here.")
		fmt.Println("[DRY RUN] System would execute 'shutdown /s /t 60'.")
		return nil
	}

	// /s = shutdown
	// /t 60 = time delay in seconds
	// /c "message" = comment
	cmd := exec.Command("shutdown", "/s", "/t", "60", "/c", "Steam Auto Shutdown: Download Complete")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute shutdown command: %w", err)
	}

	fmt.Println("Shutdown scheduled in 60 seconds. Run 'shutdown /a' to cancel.")
	return nil
}
