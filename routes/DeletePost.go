package routes

import (
	"NAS-Server-Web/configurations"
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
)

func DeleteGet(w http.ResponseWriter, r *http.Request) {
	session := sessionService.VerifySession(r)
	if session.IsNone() {
		http.Redirect(w, r, "/login-user", http.StatusUnauthorized)
		return
	}

	urlPath := mux.Vars(r)["path"]
	urlPath = filepath.Clean(urlPath)
	filePath := filepath.Join(configurations.Files, session.Unwrap().Username, urlPath)

	fileParentDirectory := filepath.Dir(filePath)

	if err := filesService.RemoveFile(filePath); err != nil {
		http.Redirect(w, r, "/files/"+fileParentDirectory, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/files/"+fileParentDirectory, http.StatusSeeOther)
}
