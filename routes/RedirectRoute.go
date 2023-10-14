package routes

import (
	"NAS-Server-Web/services/sessionService"
	"net/http"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("ftp")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	_, err = sessionService.GetSession(cookie)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/home/", http.StatusSeeOther)
}
