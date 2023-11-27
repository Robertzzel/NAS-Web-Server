package routes

import (
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func InlineFileGet(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO_InlineFileGet: called")
	cookie, err := r.Cookie("ftp")
	if err != nil {
		log.Println("INFO_InlineFileGet: no cookie")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	session, err := sessionService.GetSession(cookie)
	if err != nil {
		log.Println("INFO_InlineFileGet: no session")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	subpath := filepath.Clean(mux.Vars(r)["path"])
	filePath := filepath.Join(session.BasePath, subpath)

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Println("INFO_InlineFileGet: no file", filePath)
		return
	}

	if fileInfo.IsDir() {
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, filepath.Base(filePath)))

	if err = filesService.SendFile(filePath, session.BasePath, w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("INFO_InlineFileGet: file not sent")
	} else {
		log.Println("INFO_InlineFileGet: file sent", filePath)
		w.WriteHeader(http.StatusOK)
	}
}
