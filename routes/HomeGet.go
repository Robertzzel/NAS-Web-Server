package routes

import (
	"NAS-Server-Web/services/templateService"
	"net/http"
)

func HomeGet(w http.ResponseWriter, r *http.Request) {
	if err := templateService.GetFilesPage(w); err != nil {
		println(err.Error())
	}
}
