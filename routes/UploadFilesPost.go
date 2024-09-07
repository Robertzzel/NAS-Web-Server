package routes

import (
	"NAS-Server-Web/configurations"
	"NAS-Server-Web/utils"
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
)

func UploadFilesPost(w http.ResponseWriter, r *http.Request) {
	session := utils.VerifySession(r)
	if session.IsNone() {
		http.Redirect(w, r, "/login-user", http.StatusUnauthorized)
		return
	}

	err := r.ParseMultipartForm(128 << 20)
	if err != nil {
		return
	}

	path := mux.Vars(r)["path"]
	path = filepath.Clean(path)

	for _, fileSlice := range r.MultipartForm.File {
		fh := fileSlice[0]
		f, err := fh.Open()
		if err != nil {
			continue
		}
		dstPath := filepath.Join(configurations.Files, session.Unwrap().Username, path, fh.Filename)
		if err := utils.UploadFile(session.Unwrap().Username, dstPath, f, fh.Size); err != nil {
			_ = f.Close()
			continue
		}
		_ = f.Close()
	}

	if path == "." || path == "" || path == "/" {
		path = ""
	}
	http.Redirect(w, r, "/files/"+path, http.StatusSeeOther)
}
