package routes

import (
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"NAS-Server-Web/services/templateService"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func HomeGet(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO_HomeGet: Called ")
	cookie, err := r.Cookie("ftp")
	if err != nil {
		log.Println("INFO_HomeGet: no cookie ")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	session, err := sessionService.GetSession(cookie)
	if err != nil {
		log.Println("INFO_HomeGet: no session ")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	subpath := strings.TrimPrefix(mux.Vars(r)["path"], "/")
	path := filepath.Join(session.BasePath, subpath)

	files, err := filesService.GetFilesFromDirectory(path)
	if err != nil {
		log.Println("INFO_HomeGet: cannot get files from ", path)
		return
	}

	if err := templateService.GetFilesPage(w, files, strings.TrimPrefix(path, session.BasePath), session.Username); err != nil {
		log.Println("INFO_HomeGet: cannot make template ", err)
	}
}
