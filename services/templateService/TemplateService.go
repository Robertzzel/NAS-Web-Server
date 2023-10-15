package templateService

import (
	"NAS-Server-Web/models"
	"NAS-Server-Web/services/filesService"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"path/filepath"
)

func GetPage(w io.Writer, path string) error {
	t, err := template.ParseFiles(filepath.Clean("templates/base.html"), filepath.Clean(path))
	if err != nil {
		return err
	}

	return t.Execute(w, nil)
}

func GetLoginPage(w io.Writer) error {
	return GetPage(w, "templates/login.html")
}

func GetFilesPage(w io.Writer, files []models.FileDetails, currentPath, username string) error {
	t, err := template.ParseFiles("templates/base.html", "templates/home.html")
	if err != nil {
		return err
	}

	var sendData []byte
	if files != nil {
		sendData, err = json.Marshal(files)
		if err != nil {
			return err
		}
	} else {
		sendData = []byte("")
	}

	remaining, err := filesService.GetUserRemainingMemory(username)
	if err != nil {
		return err
	}
	used, err := filesService.GetUserUsedMemory(username)
	if err != nil {
		return err
	}

	v := struct {
		Files           string
		CurrentPath     string
		UsedMemory      string
		RemainingMemory string
		Username        string
	}{
		string(sendData),
		currentPath,
		formatFileSize(used),
		formatFileSize(remaining),
		username,
	}

	return t.Execute(w, v)
}

func formatFileSize(size int64) string {
	units := []string{"B", "KB", "MB", "GB"}
	index := 0
	for size >= 1024 && index < len(units)-1 {
		size /= 1024
		index++
	}
	return fmt.Sprintf("%.2d %s", size, units[index])
}
