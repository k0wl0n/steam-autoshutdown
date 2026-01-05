package cmd

import (
	"fmt"
	"os"

	"github.com/k0wl0n/steam-autoshutdown/internal/monitor"
	"github.com/spf13/cobra"
)

var (
	downloadThreshold int
	idleThreshold     int
	idleDuration      int
	dryRun            bool
	interfaceName     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "steamshutdown",
	Short: "Automatically shuts down your Mac when Steam downloads finish",
	Long: `Steam Auto Shutdown is a CLI tool that monitors your network traffic.
It waits for a download to start (high traffic) and then shuts down 
the system when traffic drops back to idle levels for a set duration.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Steam Auto Shutdown Monitor...")

		// Convert KB/s to Bytes/s for internal logic if needed,
		// but the monitor usually works with raw bytes or consistent units.
		// Let's pass the configuration to the monitor.

		config := monitor.Config{
			DownloadThresholdKB: downloadThreshold,
			IdleThresholdKB:     idleThreshold,
			IdleDurationSeconds: idleDuration,
			DryRun:              dryRun,
			InterfaceName:       interfaceName,
		}

		if err := monitor.Start(config); err != nil {
			fmt.Printf("Error during execution: %v\n", err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Define flags with human-friendly descriptions
	rootCmd.Flags().IntVarP(&downloadThreshold, "start-threshold", "s", 500, "Network speed (KB/s) to identify a download has started")
	rootCmd.Flags().IntVarP(&idleThreshold, "stop-threshold", "e", 100, "Network speed (KB/s) to identify a download has finished")
	rootCmd.Flags().IntVarP(&idleDuration, "idle-duration", "d", 60, "Duration (seconds) of low traffic before shutting down")
	rootCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Simulate the process without actually shutting down")
	rootCmd.Flags().StringVarP(&interfaceName, "interface", "i", "", "Specific network interface to monitor (e.g., en0). If empty, monitors all non-loopback interfaces.")
}
