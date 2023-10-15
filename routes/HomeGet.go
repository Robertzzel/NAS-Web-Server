package routes

import (
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"NAS-Server-Web/services/templateService"
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
	"strings"
)

func HomeGet(w http.ResponseWriter, r *http.Request) {
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

	subpath := strings.TrimPrefix(mux.Vars(r)["path"], "/")
	path := filepath.Join(session.BasePath, subpath)

	files, err := filesService.GetFilesFromDirectory(path)
	if err != nil {
		return
	}

	if err := templateService.GetFilesPage(w, files, strings.TrimPrefix(path, session.BasePath), session.Username); err != nil {
		println(err.Error())
	}
}
