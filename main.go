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

	r.HandleFunc("/cd", routes.CdPost).Methods("POST")
	r.HandleFunc("/cd", routes.Redirect).Methods("GET")

	r.HandleFunc("/cdup", routes.CdupPost).Methods("POST")
	r.HandleFunc("/cdup", routes.Redirect).Methods("GET")

	//r.HandleFunc("/api/list", ListPost).Methods("POST")
	//r.HandleFunc("/api/rm", RemovePost).Methods("POST")
	//r.HandleFunc("/api/dwat/{file}", DownloadFileAttachmentGet).Methods("GET")
	//r.HandleFunc("/api/dwin/{file}", DownloadFileInlineGet).Methods("GET")
	//r.HandleFunc("/api/dwdr/{file}", DownloadDirectoryGet).Methods("GET")
	//r.HandleFunc("/api/upload/{name}", UploadFilesPost).Methods("POST")
	//r.HandleFunc("/api/directory", CreateDirectoryPost).Methods("POST")
	//r.HandleFunc("/api/rename", RenameFilePost).Methods("POST")
	//r.HandleFunc("/api/details", UserDetailsGet).Methods("GET")

	fmt.Println("Starting on " + configs.GetHost() + ":" + configs.GetPort())
	log.Fatal(http.ListenAndServe(configs.GetHost()+":"+configs.GetPort(), r))
}
