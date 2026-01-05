package monitor

import (
	"testing"
)

func TestConfig(t *testing.T) {
	// Simple test to verify Config struct structure
	cfg := Config{
		DownloadThresholdKB: 500,
		IdleThresholdKB:     100,
		IdleDurationSeconds: 60,
		DryRun:              true,
		InterfaceName:       "eth0",
	}

	if cfg.DownloadThresholdKB != 500 {
		t.Errorf("Expected DownloadThresholdKB to be 500, got %d", cfg.DownloadThresholdKB)
	}
	if !cfg.DryRun {
		t.Error("Expected DryRun to be true")
	}
}
