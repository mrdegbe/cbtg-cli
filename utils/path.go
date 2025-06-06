package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// ValidatePath checks if the given path exists and returns the absolute version
func ValidatePath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("invalid path: %w", err)
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return "", fmt.Errorf("path does not exist: %s", absPath)
	}

	return absPath, nil
}
