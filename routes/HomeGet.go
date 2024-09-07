package routes

import (
	"NAS-Server-Web/configurations"
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"NAS-Server-Web/services/templates"
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
)

func HomeGet(w http.ResponseWriter, r *http.Request) {
	session := sessionService.VerifySession(r)
	if session.IsNone() {
		http.Redirect(w, r, "/login-user", http.StatusUnauthorized)
		return
	}

	path, exists := mux.Vars(r)["path"]
	if !exists {
		http.Redirect(w, r, "/files/", http.StatusUnauthorized)
		return
	}
	path = filepath.Clean(path)
	path = filepath.Join(configurations.Files, session.Unwrap().Username, path)

	files := filesService.GetFilesFromDirectory(path)
	if files.IsError() {
		return
	}

	_ = templates.WriteFilesPage(w, files.Unwrap(), session.Unwrap().Username)
}
