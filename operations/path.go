package operations

import (
	"os"
	"path/filepath"
	"strings"
)

func IsPathSafe(path string) bool {
	return !strings.Contains(path, ".")
}

func DirSize(path string) int64 {
	var dirSize int64 = 0

	readSize := func(path string, file os.FileInfo, err error) error {
		if file != nil && !file.IsDir() {
			dirSize += file.Size()
		}

		return nil
	}

	filepath.Walk(path, readSize)

	return dirSize
}
