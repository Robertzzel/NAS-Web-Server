package routes

import (
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
)

func InlineFileGet(w http.ResponseWriter, r *http.Request) {
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

	subpath := filepath.Clean(mux.Vars(r)["path"])
	filePath := filepath.Join(session.BasePath, subpath)

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return
	}

	if fileInfo.IsDir() {
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, filepath.Base(filePath)))

	if err = filesService.SendFile(filePath, w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
