package operations

import (
	. "NAS-Server-Web/settings"
	"errors"
	"os"
)

func GetUserUsedMemory(username string) (int64, error) {
	entries, err := os.ReadDir(BasePath)
	if err != nil {
		return 0, err
	}

	for _, dir := range entries {
		if dir.Name() != username {
			continue
		}
		info, err := dir.Info()
		if err != nil {
			return 0, err
		}
		dirSize := DirSize(BasePath + "/" + info.Name())
		return dirSize, nil
	}

	return 0, errors.New("username does not exist")
}

func GetUserRemainingMemory(username string) (int64, error) {
	used, err := GetUserUsedMemory(username)
	if err != nil {
		return 0, err
	}
	return MemoryPerUsed - used, nil
}
