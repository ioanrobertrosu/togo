package utils

import (
	"fmt"
	"os"
)

func EnsureFileExists(filename string, truncate bool) (*os.File, error) {
	flags := os.O_CREATE | os.O_WRONLY

	if truncate {
		flags |= os.O_TRUNC
	}

	file, err := os.OpenFile(filename, flags, 0644)

	if err != nil {
		return nil, fmt.Errorf("couldn't open or create '%s': %w", filename, err)
	}

	return file, nil
}