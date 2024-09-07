package routes

import (
	"NAS-Server-Web/configurations"
	"NAS-Server-Web/services/sessionService"
	"net/http"
	"os"
	"path/filepath"
)

func RenamePost(w http.ResponseWriter, r *http.Request) {
	session := sessionService.VerifySession(r)
	if session.IsNone() {
		http.Redirect(w, r, "/login-user", http.StatusUnauthorized)
		return
	}

	newPath := r.URL.Query().Get("new_path")
	oldPath := r.URL.Query().Get("old_path")
	if newPath == "" || oldPath == "" {
		http.Redirect(w, r, "/files/", http.StatusUnauthorized)
		return
	}

	newPath = filepath.Clean(newPath)
	oldPath = filepath.Clean(oldPath)

	fullOldPath := filepath.Join(configurations.Files, session.Unwrap().Username, oldPath)
	fullNewPath := filepath.Join(configurations.Files, session.Unwrap().Username, newPath)

	_ = os.Rename(fullOldPath, fullNewPath)

	http.Redirect(w, r, "/files/"+filepath.Dir(newPath), http.StatusSeeOther)
}
