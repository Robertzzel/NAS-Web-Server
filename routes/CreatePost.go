package routes

import (
	"NAS-Server-Web/services/sessionService"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO_CreatePost: Create dir called")
	cookie, err := r.Cookie("ftp")
	if err != nil {
		log.Println("INFO_CreatePost: No cookie redirecting")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	session, err := sessionService.GetSession(cookie)
	if err != nil {
		log.Println("INFO_CreatePost: No good session redirecting")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		return
	}

	var data map[string]string
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println("INFO_CreatePost: Bad request, failed to decode json")
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
		log.Println("INFO_CreatePost: bad path")
		http.Redirect(w, r, "/home/"+currentPath, http.StatusSeeOther)
		return
	}

	log.Println("INFO_CreatePost: Created", "/home/"+currentPath)
	http.Redirect(w, r, "/home/"+currentPath, http.StatusSeeOther)
}
