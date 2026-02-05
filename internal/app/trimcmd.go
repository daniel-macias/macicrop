package app

import (
	"flag"
	"fmt"
	"os"
)

type TrimOptions struct {
	InputDir  string
	OutputDir string

	Overwrite bool
	Suffix    string
	KeepEmpty bool
}

func ParseTrimArgs(args []string) (TrimOptions, error) {
	fs := flag.NewFlagSet("trim", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)

	var opts TrimOptions
	fs.BoolVar(&opts.Overwrite, "overwrite", false, "overwrite existing files in output directory")
	fs.StringVar(&opts.Suffix, "suffix", "", "append suffix to output filenames (e.g. _trim)")
	fs.BoolVar(&opts.KeepEmpty, "keep-empty", true, "write 1x1 transparent PNG for fully transparent images")

	if err := fs.Parse(args); err != nil {
		return TrimOptions{}, err
	}

	rest := fs.Args()
	if len(rest) != 2 {
		fmt.Println("Usage: macicrop trim <input_dir> <output_dir> [--overwrite] [--suffix _trim] [--keep-empty]")
		return TrimOptions{}, fmt.Errorf("expected input_dir and output_dir")
	}

	opts.InputDir = rest[0]
	opts.OutputDir = rest[1]
	return opts, nil
}
