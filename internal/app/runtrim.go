package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type TrimStats struct {
	FoundPNGs int
}

func RunTrim(opts TrimOptions) (TrimStats, error) {
	// Validate input dir exists
	info, err := os.Stat(opts.InputDir)
	if err != nil {
		return TrimStats{}, fmt.Errorf("input dir error: %w", err)
	}
	if !info.IsDir() {
		return TrimStats{}, fmt.Errorf("input path is not a directory: %s", opts.InputDir)
	}

	// Ensure output dir exists
	if err := os.MkdirAll(opts.OutputDir, 0o755); err != nil {
		return TrimStats{}, fmt.Errorf("failed to create output dir: %w", err)
	}

	entries, err := os.ReadDir(opts.InputDir)
	if err != nil {
		return TrimStats{}, fmt.Errorf("failed to read input dir: %w", err)
	}

	stats := TrimStats{}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.ToLower(filepath.Ext(name)) == ".png" {
			stats.FoundPNGs++
		}
	}

	return stats, nil
}
