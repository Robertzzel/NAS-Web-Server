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
		return
	}

	session, err := sessionService.GetSession(cookie)
	if err != nil {
		return
	}

	subpath := mux.Vars(r)["path"]
	filePath := filepath.Join(session.BasePath, subpath)

	if err = filesService.RemoveFile(filePath); err != nil {
		return
	}

	fileParentDirectory := filepath.Dir(subpath)
	if fileParentDirectory == "." {
		fileParentDirectory = ""
	}
	http.Redirect(w, r, "/home/"+fileParentDirectory, http.StatusSeeOther)
}
