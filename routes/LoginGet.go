package routes

import (
	"NAS-Server-Web/utils"
	"net/http"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {
	_ = utils.WriteLoginPage(w)
}
