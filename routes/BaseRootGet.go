package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func BaseRootGet(w http.ResponseWriter, r *http.Request) {
	path, pathExists := mux.Vars(r)["file"]
	if !pathExists {
		return
	}

	if strings.HasSuffix(path, "js") {
		w.Header().Set("Content-Type", "text/javascript")
	} else {
		w.Header().Set("Content-Type", "text/css")
	}

	file := filepath.Join("./static", filepath.Clean(path))
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	_, err = w.Write(data)
	if err != nil {
		fmt.Print(err)
	}
}
