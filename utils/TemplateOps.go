package utils

import (
	"NAS-Server-Web/models"
	"encoding/json"
	"html/template"
	"io"
)

func WritePage(w io.Writer, filePath string) error {
	t, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}

	return t.Execute(w, nil)
}

func WriteLoginPage(w io.Writer) error {
	return WritePage(w, "templates/login.html")
}

func WriteFilesPage(w io.Writer, files []models.FileDetails, username string) error {
	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		return err
	}

	sendData, err := json.Marshal(files)
	if err != nil {
		return err
	}

	v := struct {
		Files    string
		Username string
	}{
		string(sendData),
		username,
	}

	return t.Execute(w, v)
}
