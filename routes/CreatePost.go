package routes

import (
	"NAS-Server-Web/configurations"
	"NAS-Server-Web/services/sessionService"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	sessionCookie := sessionService.VerifySession(r)
	if sessionCookie.IsNone() {
		http.Redirect(w, r, "/login-user", http.StatusUnauthorized)
		return
	}

	path := mux.Vars(r)["path"]
	path = filepath.Clean(path)

	err := os.Mkdir(filepath.Join(configurations.Files, sessionCookie.Unwrap().Username, path), 0770)
	if err != nil {
		http.Redirect(w, r, "/files/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, filepath.Join("/files/", path), http.StatusSeeOther)
}
