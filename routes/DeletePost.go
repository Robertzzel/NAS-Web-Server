package routes

import (
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
)

func DeleteGet(w http.ResponseWriter, r *http.Request) {
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

	subpath := mux.Vars(r)["path"]
	filePath := filepath.Join(session.BasePath, subpath)

	fileParentDirectory := filepath.Dir(subpath)
	if fileParentDirectory == "." {
		fileParentDirectory = ""
	}

	if err = filesService.RemoveFile(filePath); err != nil {
		http.Redirect(w, r, "/home/"+fileParentDirectory, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/home/"+fileParentDirectory, http.StatusSeeOther)
}
