package routes

import (
	"NAS-Server-Web/utils"
	"net/http"
)

func LoginPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		return
	}

	if !r.PostForm.Has("username") || !r.PostForm.Has("password") {
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	credentialsCheckResult := utils.CheckUsernameAndPassword(username, password)
	if credentialsCheckResult.IsError() || !credentialsCheckResult.Unwrap() {
		http.Redirect(w, r, "/login-user", http.StatusSeeOther)
		return
	}

	cookie := utils.CreateSession(username)
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/files/", http.StatusSeeOther)
}
