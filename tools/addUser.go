package main

import (
	"NAS-Server-Web/services/configsService"
	"NAS-Server-Web/services/databaseService"
	"errors"
	"os"
	"path"
)

func main() {
	if len(os.Args) != 4 {
		println("must give Username, Email and Password")
		os.Exit(1)
	}

	configs, err := configsService.NewConfigsService()
	if err != nil {
		panic(err)
	}

	db, err := databaseService.NewDatabaseService()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.AddUser(os.Args[1], os.Args[2], os.Args[3]); err != nil {
		println("cannot add user ", err.Error())
		os.Exit(1)
	}

	err = os.Mkdir(configs.GetBaseFilesPath(), 0777)
	if err != nil && !errors.Is(err, os.ErrExist) {
		panic(err)
	}

	fullPath := path.Join(configs.GetBaseFilesPath(), os.Args[1])
	if err = os.Mkdir(fullPath, 0777); err != nil {
		panic(err)
	}
}
