package routes

import (
	"NAS-Server-Web/services/sessionService"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
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

	if err := r.ParseForm(); err != nil {
		return
	}

	var data map[string]string
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	name, hasNewName := data["name"]
	path, hasOldPath := data["path"]
	if !hasOldPath || !hasNewName {
		return
	}

	currentPath := filepath.Clean(path)
	directoryName := filepath.Clean(name)
	directoryName = filepath.Join(currentPath, directoryName)

	if currentPath == "." || currentPath == "/" {
		currentPath = ""
	}

	err = os.Mkdir(filepath.Join(session.BasePath, directoryName), 0770)
	if err != nil {
		http.Redirect(w, r, "/home/"+currentPath, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/home/"+currentPath, http.StatusSeeOther)
}
