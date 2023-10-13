package templateService

import (
	"NAS-Server-Web/models"
	"encoding/json"
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

func GetFilesPage(w io.Writer, files []models.FileDetails) error {
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

	return t.Execute(w, string(sendData))
}
