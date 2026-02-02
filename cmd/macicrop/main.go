package main

import (
	"fmt"
	"os"

	"github.com/daniel-macias/macicrop/internal/app"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(2)
	}

	switch os.Args[1] {
	case "trim":
		opts, err := app.ParseTrimArgs(os.Args[2:])
		if err != nil {
			os.Exit(2)
		}
		fmt.Printf("opts: %+v\n", opts)
		stats, err := app.RunTrim(opts)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		fmt.Printf("found %d pngs\n", stats.FoundPNGs)
	default:
		fmt.Printf("unknown command: %s\n\n", os.Args[1])
		printUsage()
		os.Exit(2)
	}
}

func printUsage() {
	fmt.Println(`macicrop - crop PNGs to visible (non-transparent) bounds

Usage:
  macicrop trim <input_dir> <output_dir> [--overwrite] [--suffix _trim] [--keep-empty]
`)
}
