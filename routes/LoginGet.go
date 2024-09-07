package routes

import (
	"NAS-Server-Web/services/templates"
	"net/http"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {
	_ = templates.WriteLoginPage(w)
}
