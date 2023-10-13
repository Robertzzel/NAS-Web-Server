package routes

import (
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"NAS-Server-Web/services/templateService"
	"net/http"
	"path/filepath"
)

func HomeGet(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("ftp")
	if err != nil {
		return
	}

	session, err := sessionService.GetSession(cookie)
	if err != nil {
		return
	}

	if err := r.ParseForm(); err != nil {
		return
	}

	var path string
	if !r.PostForm.Has("path") {
		path = session.BasePath
	} else {
		path = filepath.Join(session.BasePath, filepath.Clean(r.FormValue("path")))
	}

	files, err := filesService.GetFilesFromDirectory(path)
	if err != nil {
		return
	}

	if err := templateService.GetFilesPage(w, files); err != nil {
		println(err.Error())
	}
}
