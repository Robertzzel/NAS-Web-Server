package main

import (
	"NAS-Server-Web/configurations"
	"NAS-Server-Web/routes"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	blockedIps sync.Map
)

func BlockIP(r *http.Request) {
	ip := r.RemoteAddr // Get IP address of the requester

	blockInfo, isBlocked := blockedIps.Load(ip)

	if isBlocked {
		if blockInfo.(time.Time).After(
			time.Now()) {
			return
		}
		blockedIps.Delete(ip)
	}

	if r.URL.Path == "/" || r.URL.Path == "/login" {
		until := time.Now().Add(time.Minute)
		blockedIps.Store(ip, until)
		log.Println(ip, "blocked until", until)
		return
	}
}

func main() {
	if len(os.Args) != 7 {
		log.Fatal("ex: ./server <host> <port> <database> <files> <ssl_cert> <ssl_key>")
	}
	configurations.Host = os.Args[1]
	configurations.Port = os.Args[2]
	configurations.Database = os.Args[3]
	configurations.Files = os.Args[4]
	configurations.SslCertificatePath = os.Args[5]
	configurations.SslKeyPath = os.Args[6]

	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { BlockIP(r) })

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.HandleFunc("/login-user", routes.LoginGet).Methods("GET")
	r.HandleFunc("/login-user", routes.LoginPost).Methods("POST")

	r.HandleFunc("/files/{path:.*}", routes.HomeGet).Methods("GET")
	r.HandleFunc("/delete/{path:.*}", routes.DeleteGet).Methods("GET")
	r.HandleFunc("/download/{path:.*}", routes.DownloadGet).Methods("GET")
	r.HandleFunc("/inline/{path:.*}", routes.InlineFileGet).Methods("GET")
	r.HandleFunc("/rename", routes.RenamePost).Methods("GET")
	r.HandleFunc("/create/{path:.*}", routes.CreatePost).Methods("GET")
	r.HandleFunc("/upload/{path:.*}", routes.UploadFilesPost).Methods("POST")

	fmt.Println("Starting on " + configurations.Host + ":" + configurations.Port)
	err := http.ListenAndServeTLS(configurations.Host+":"+configurations.Port, configurations.SslCertificatePath, configurations.SslKeyPath, r)
	log.Fatal(err)
}
