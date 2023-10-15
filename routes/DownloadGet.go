package routes

import (
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
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

	subpath := mux.Vars(r)["path"]
	filePath := filepath.Join(session.BasePath, subpath)

	fileToSend, err := filesService.PrepareFile(filePath)
	if err != nil {
		return
	}

	fileInfo, err := os.Stat(fileToSend)
	if err != nil {
		return
	}

	setHeaders(w, filepath.Base(filePath)+".zip", strconv.Itoa(int(fileInfo.Size())))
	w.WriteHeader(http.StatusOK)

	fileHandler, err := os.Open(fileToSend)
	if err != nil {
		return
	}
	defer fileHandler.Close()

	_, err = io.Copy(w, fileHandler)
	if err != nil {
		return
	}

	originalFileStat, err := os.Stat(filePath)
	if err != nil {
		return
	}

	if originalFileStat.IsDir() {
		err := os.Remove(fileToSend)
		if err != nil {
			log.Println("Cannot delete remnant on zipping", filePath, "zip name", fileToSend, err.Error())
		}
	}
}

func setHeaders(w http.ResponseWriter, name, len string) {
	//Represents binary file
	w.Header().Set("Content-Type", "application/octet-stream")
	//Tells client what filename should be used.
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, name))
	//The length of the data.
	w.Header().Set("Content-Length", len)
	//No cache headers.
	w.Header().Set("Cache-Control", "private")
	//No cache headers.
	w.Header().Set("Pragma", "private")
	//No cache headers.
	w.Header().Set("Expires", "Mon, 26 Jul 1997 05:00:00 GMT")
}
