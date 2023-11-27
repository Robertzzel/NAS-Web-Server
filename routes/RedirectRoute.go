package routes

import (
	"NAS-Server-Web/services/sessionService"
	"log"
	"net/http"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO_Redirect: called")
	cookie, err := r.Cookie("ftp")
	if err != nil {
		log.Println("INFO_Redirect: no cookie")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	_, err = sessionService.GetSession(cookie)
	if err != nil {
		log.Println("INFO_Redirect: no session")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	log.Println("INFO_Redirect: redirected")
	http.Redirect(w, r, "/home/", http.StatusSeeOther)
}
