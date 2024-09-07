package databaseService

import (
	"NAS-Server-Web/configurations"
	. "NAS-Server-Web/utils"
	"crypto/sha256"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var instance *sql.DB = nil

func getDatabase() Result[*sql.DB] {
	if instance == nil {
		db, err := sql.Open("sqlite3", configurations.Database)
		if err != nil {
			return Error[*sql.DB](err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS User(
			Id integer PRIMARY KEY,
			Name varchar(255) UNIQUE NOT NULL,
			Email varchar(255),
			Password varchar(255) NOT NULL
		)`)
		if err != nil {
			return Error[*sql.DB](err)
		}

		instance = db
	}

	return Ok(instance)
}

func CheckUsernameAndPassword(username, password string) Result[bool] {
	databaseResult := getDatabase()
	db := databaseResult.Unwrap()

	var cnt int
	err := db.QueryRow(`select count(*) from User where Name = ? and Password = ? LIMIT 1`, username, hash(password)).Scan(&cnt)
	if err != nil {
		return Error[bool](err)
	}

	return Ok(cnt != 0)
}

func AddUser(username, email, password string) error {
	databaseResult := getDatabase()
	db := databaseResult.Unwrap()

	_, err := db.Exec(`INSERT INTO User (Name, Email, PASSWORD) VALUES (?, ?, ?)`, username, email, hash(password))
	return err
}

func hash(password string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}
