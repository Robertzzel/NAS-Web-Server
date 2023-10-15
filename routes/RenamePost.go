package routes

import (
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"net/http"
	"path/filepath"
)

func RenamePost(w http.ResponseWriter, r *http.Request) {
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

	if !r.PostForm.Has("new-name") || !r.PostForm.Has("old-path") {
		return
	}

	newName := filepath.Clean(r.FormValue("new-name"))
	oldPath := filepath.Clean(r.FormValue("old-path"))

	newName = filepath.Join(filepath.Dir(oldPath), newName)
	fileDirectory := filepath.Dir(oldPath)
	if fileDirectory == "." || fileDirectory == "/" {
		fileDirectory = ""
	}

	fullOldPath := filepath.Join(session.BasePath, oldPath)
	fullNewPath := filepath.Join(session.BasePath, newName)
	if err = filesService.RenameFile(fullOldPath, fullNewPath); err != nil {
		http.Redirect(w, r, "/home/"+fileDirectory, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/home/"+fileDirectory, http.StatusSeeOther)
}
