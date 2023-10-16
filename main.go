package main

import (
	"NAS-Server-Web/routes"
	"NAS-Server-Web/services/configsService"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	configs, err := configsService.NewConfigsService()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", routes.Redirect).Methods("GET")

	r.HandleFunc("/login", routes.LoginGet).Methods("GET")
	r.HandleFunc("/login", routes.LoginPost).Methods("POST")

	r.HandleFunc("/home/{path:.*}", routes.HomeGet).Methods("GET")

	r.HandleFunc("/delete/{path:.*}", routes.DeleteGet).Methods("GET")

	r.HandleFunc("/file/{path:.*}", routes.DownloadGet).Methods("GET")

	r.HandleFunc("/inline/{path:.*}", routes.InlineFileGet).Methods("GET")

	r.HandleFunc("/rename", routes.RenamePost).Methods("POST")

	r.HandleFunc("/create", routes.CreatePost).Methods("POST")

	r.HandleFunc("/upload", routes.UploadFilesPost).Methods("POST")

	fmt.Println("Starting on " + configs.GetHost() + ":" + configs.GetPort())
	log.Fatal(http.ListenAndServeTLS(configs.GetHost()+":"+configs.GetPort(), configs.GetCertificateFilePath(), configs.GetKeyFilePath(), r))
}
