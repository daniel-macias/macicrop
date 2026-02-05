package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func IsPNG(name string) bool {
	return strings.ToLower(filepath.Ext(name)) == ".png"
}

func OutputName(filename, suffix string) (string, error) {
	ext := filepath.Ext(filename)
	if strings.ToLower(ext) != ".png" {
		return "", fmt.Errorf("not a png filename: %s", filename)
	}

	base := strings.TrimSuffix(filename, ext)
	if suffix == "" {
		return filename, nil
	}
	return base + suffix + ext, nil
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
