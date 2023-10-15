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
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	session, err := sessionService.GetSession(cookie)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
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

	if currentPath == "." || currentPath == "/" {
		currentPath = ""
	}

	err = os.Mkdir(filepath.Join(session.BasePath, directoryName), 0770)
	if err != nil {
		http.Redirect(w, r, "/home/"+currentPath, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/home/"+currentPath, http.StatusSeeOther)
}
