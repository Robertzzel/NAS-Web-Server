package routes

import (
	"NAS-Server-Web/services/templateService"
	"log"
	"net/http"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO_LoginGet: called")
	if err := templateService.GetLoginPage(w); err != nil {
		log.Println("INFO_LoginGet: error loading page")
		println(err.Error())
	}
	log.Println("INFO_LoginGet: page loaded")
}
