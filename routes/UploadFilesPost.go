package routes

import (
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"log"
	"net/http"
	"path/filepath"
)

func UploadFilesPost(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO_UploadFilesPost: called")
	cookie, err := r.Cookie("ftp")
	if err != nil {
		log.Println("INFO_UploadFilesPost: no cookie")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	session, err := sessionService.GetSession(cookie)
	if err != nil {
		log.Println("INFO_UploadFilesPost: no session")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	err = r.ParseMultipartForm(32 << 20) // 32 MB
	if err != nil {
		log.Println("INFO_UploadFilesPost: cannt parse form", err)
		return
	}

	path := r.FormValue("path")
	path = filepath.Clean(path)

	for _, fh := range r.MultipartForm.File["files"] {
		f, err := fh.Open()
		if err != nil {
			continue
		}
		dstPath := filepath.Join(session.BasePath, path, fh.Filename)
		if err := filesService.UploadFile(session.Username, dstPath, f, fh.Size); err != nil {
			_ = f.Close()
			log.Println("INFO_UploadFilesPost: cannot upload file")
			continue
		}
		_ = f.Close()
	}

	if path == "." || path == "" || path == "/" {
		path = ""
	}
	http.Redirect(w, r, "/home/"+path, http.StatusSeeOther)
}
