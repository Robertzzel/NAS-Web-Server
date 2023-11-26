package routes

import (
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func DownloadGet(w http.ResponseWriter, r *http.Request) {
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
		setHeaders(w, filepath.Base(filePath)+".zip", "")
	} else {
		setHeaders(w, filepath.Base(filePath), strconv.Itoa(int(fileInfo.Size())))
	}

	if err = filesService.SendFile(filePath, session.BasePath, w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}

func setHeaders(w http.ResponseWriter, name, len string) {
	//Represents binary file
	w.Header().Set("Content-Type", "application/octet-stream")
	//Tells client what filename should be used.
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, name))
	//The length of the data.
	if len != "" {
		w.Header().Set("Content-Length", len)
	}
	//No cache headers.
	w.Header().Set("Cache-Control", "private")
	//No cache headers.
	w.Header().Set("Pragma", "private")
	//No cache headers.
	w.Header().Set("Expires", "Mon, 26 Jul 1997 05:00:00 GMT")
}
