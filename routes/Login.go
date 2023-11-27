package routes

import (
	"NAS-Server-Web/services/databaseService"
	"NAS-Server-Web/services/sessionService"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func LoginPost(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO_LoginPost: called")
	if err := r.ParseForm(); err != nil {
		return
	}

	if !r.PostForm.Has("username") || !r.PostForm.Has("password") {
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	log.Println("INFO_LoginPost: login params", username, password)
	dbInstance, err := databaseService.NewDatabaseService()
	if err != nil {
		log.Println("INFO_LoginPost: cannot get the daatbase")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	ok, err := dbInstance.UsernameAndPasswordExists(username, password)
	if err != nil {
		log.Println("INFO_LoginPost: error verifying the db")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if !ok {
		log.Println("INFO_LoginPost: wrong credentials")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	cookie := new(http.Cookie)
	cookie.Name = "ftp"
	cookie.Value = uuid.New().String()
	cookie.Expires = time.Now().Add(24 * time.Hour)

	if err = sessionService.NewSession(cookie.Value, username); err != nil {
		log.Println("INFO_LoginPost: session not created")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	log.Println("INFO_LoginPost: session created")
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/home/", http.StatusSeeOther)
}
