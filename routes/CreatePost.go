package routes

import (
	"NAS-Server-Web/services/sessionService"
	"net/http"
	"os"
	"path/filepath"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
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

	if !r.PostForm.Has("path") || !r.PostForm.Has("name") {
		return
	}

	currentPath := filepath.Clean(r.FormValue("path"))
	directoryName := filepath.Clean(r.FormValue("name"))
	directoryName = filepath.Join(currentPath, directoryName)

	err = os.Mkdir(filepath.Join(session.BasePath, directoryName), 0770)
	if err != nil {
		return
	}

	if currentPath == "." || currentPath == "/" {
		currentPath = ""
	}
	http.Redirect(w, r, "/home/"+currentPath, http.StatusSeeOther)
}
