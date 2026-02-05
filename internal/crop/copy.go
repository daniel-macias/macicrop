package crop

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(inPath, outPath string) error {
	in, err := os.Open(inPath)
	if err != nil {
		return fmt.Errorf("open input: %w", err)
	}
	defer in.Close()

	out, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("create output: %w", err)
	}

	_, copyErr := io.Copy(out, in)
	closeErr := out.Close()

	if copyErr != nil {
		return fmt.Errorf("copy: %w", copyErr)
	}
	if closeErr != nil {
		return fmt.Errorf("close output: %w", closeErr)
	}

	return nil
}
