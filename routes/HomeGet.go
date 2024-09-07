package routes

import (
	"NAS-Server-Web/configurations"
	"NAS-Server-Web/utils"
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
)

func HomeGet(w http.ResponseWriter, r *http.Request) {
	session := utils.VerifySession(r)
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

	files := utils.GetFilesFromDirectory(path)
	if files.IsError() {
		return
	}

	_ = utils.WriteFilesPage(w, files.Unwrap(), session.Unwrap().Username)
}
