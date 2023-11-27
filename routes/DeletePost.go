package routes

import (
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path/filepath"
)

func DeleteGet(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO_DeleteGet: Delete called")
	cookie, err := r.Cookie("ftp")
	if err != nil {
		log.Println("INFO_DeleteGet: no cookei redirecting")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	session, err := sessionService.GetSession(cookie)
	if err != nil {
		log.Println("INFO_DeleteGet: no good session redirecting")
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
		log.Println("INFO_DeleteGet: cannot remove file", filePath)
		http.Redirect(w, r, "/home/"+fileParentDirectory, http.StatusSeeOther)
		return
	}

	log.Println("INFO_DeleteGet: Removed", filePath)
	http.Redirect(w, r, "/home/"+fileParentDirectory, http.StatusSeeOther)
}
