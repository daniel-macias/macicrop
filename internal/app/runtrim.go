package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/daniel-macias/macicrop/internal/crop"
	"github.com/daniel-macias/macicrop/internal/fs"
)

type TrimStats struct {
	Found   int
	Trimmed int
	Empty   int
	Skipped int
	Errors  int
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
		if !fs.IsPNG(name) {
			continue
		}

		stats.Found++

		outName, err := fs.OutputName(name, opts.Suffix)
		if err != nil {
			// should never happen maybe TODO test this is for safety
			stats.Errors++
			fmt.Printf("error   %s (%v)\n", name, err)
			continue
		}

		inPath := filepath.Join(opts.InputDir, name)
		outPath := filepath.Join(opts.OutputDir, outName)

		// Collision handling
		exists, err := fs.Exists(outPath)
		if err != nil {
			stats.Errors++
			fmt.Printf("error   %s (stat output: %v)\n", name, err)
			continue
		}

		if exists && !opts.Overwrite {
			stats.Skipped++
			fmt.Printf("skip    %s (exists)\n", outName)
			continue
		}

		// Testing copying the files only
		if err := crop.CopyFile(inPath, outPath); err != nil {
			stats.Errors++
			fmt.Printf("error   %s (%v)\n", name, err)
			continue
		}

		stats.Trimmed++
		fmt.Printf("copied  %s -> %s\n", name, outName)
	}

	return stats, nil
}
