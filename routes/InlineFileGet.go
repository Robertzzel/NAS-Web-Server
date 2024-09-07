package routes

import (
	"NAS-Server-Web/configurations"
	"NAS-Server-Web/utils"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func InlineFileGet(w http.ResponseWriter, r *http.Request) {
	session := utils.VerifySession(r)
	if session.IsNone() {
		http.Redirect(w, r, "/login-user", http.StatusUnauthorized)
		return
	}

	urlPath := mux.Vars(r)["path"]
	urlPath = filepath.Clean(urlPath)
	filePath := filepath.Join(configurations.Files, session.Unwrap().Username, urlPath)

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Println("INFO_InlineFileGet: no file", filePath)
		return
	}

	if fileInfo.IsDir() {
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, filepath.Base(filePath)))

	if err = utils.SendFile(filePath, w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
