package routes

import (
	"NAS-Server-Web/services/filesService"
	"NAS-Server-Web/services/sessionService"
	"net/http"
	"path/filepath"
)

func UploadFilesPost(w http.ResponseWriter, r *http.Request) {
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

	path := r.FormValue("path")
	path = filepath.Clean(path)

	err = r.ParseMultipartForm(32 << 20) // 32 MB
	if err != nil {
		return
	}

	for _, fh := range r.MultipartForm.File["files"] {
		f, err := fh.Open()
		if err != nil {
			continue
		}
		dstPath := filepath.Join(session.BasePath, path, fh.Filename)
		if err := filesService.UploadFile(session.Username, dstPath, f, fh.Size); err != nil {
			_ = f.Close()
			continue
		}
		_ = f.Close()
	}

	http.Redirect(w, r, "/home/"+path, http.StatusSeeOther)
}
