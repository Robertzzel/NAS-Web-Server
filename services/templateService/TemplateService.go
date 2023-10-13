package templateService

import (
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

func GetFilesPage(w io.Writer) error {
	return GetPage(w, "templates/home.html")
}
