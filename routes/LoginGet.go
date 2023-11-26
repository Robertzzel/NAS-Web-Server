package routes

import (
	"NAS-Server-Web/services/templateService"
	"net/http"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {
	if err := templateService.GetLoginPage(w); err != nil {
		println(err.Error())
	}
}
