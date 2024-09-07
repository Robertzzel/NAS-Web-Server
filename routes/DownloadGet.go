package routes

import (
	"NAS-Server-Web/configurations"
	"NAS-Server-Web/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func DownloadGet(w http.ResponseWriter, r *http.Request) {
	session := utils.VerifySession(r)
	if session.IsNone() {
		http.Redirect(w, r, "/login-user", http.StatusUnauthorized)
		return
	}

	urlPath := mux.Vars(r)["path"]
	filePath := filepath.Join(configurations.Files, session.Unwrap().Username, urlPath)

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		http.Redirect(w, r, "/files/", http.StatusUnauthorized)
		return
	}

	if fileInfo.IsDir() {
		setHeaders(w, filepath.Base(filePath)+".zip", "")
	} else {
		setHeaders(w, filepath.Base(filePath), strconv.Itoa(int(fileInfo.Size())))
	}

	if err = utils.SendFile(filePath, w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}

func setHeaders(w http.ResponseWriter, name, len string) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, name))
	if len != "" {
		w.Header().Set("Content-Length", len)
	}
	w.Header().Set("Cache-Control", "private")
	w.Header().Set("Pragma", "private")
}
