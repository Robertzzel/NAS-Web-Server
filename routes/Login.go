package routes

import (
	"NAS-Server-Web/services/databaseService"
	"NAS-Server-Web/services/sessionService"
	"NAS-Server-Web/services/templateService"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {
	if err := templateService.GetLoginPage(w); err != nil {
		println(err.Error())
	}
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if !r.PostForm.Has("username") || !r.PostForm.Has("password") {
		println("BAD")
		return
	}
	println("GOOD")
	username := r.FormValue("username")
	password := r.FormValue("password")

	dbInstance, err := databaseService.NewDatabaseService()
	if err != nil {
		return
	}
	ok, err := dbInstance.Login(username, password)
	if err != nil {
		return
	}
	if !ok {
		return
	}

	cookie := new(http.Cookie)
	cookie.Name = "ftp"
	cookie.Value = uuid.New().String()
	cookie.Expires = time.Now().Add(24 * time.Hour)

	if err := sessionService.NewSession(cookie.Value, username); err != nil {
		return
	}

	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
