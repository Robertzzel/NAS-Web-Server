package databaseService

import (
	"NAS-Server-Web/services/configsService"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseService struct {
	*sql.DB
}

var instance *DatabaseService = nil

func NewDatabaseService() (*DatabaseService, error) {
	if instance == nil {
		configs, err := configsService.NewConfigsService()
		if err != nil {
			return nil, err
		}

		db, err := sql.Open("sqlite3", configs.GetDatabasePath())
		if err != nil {
			return nil, err
		}

		dm := DatabaseService{db}
		if err = dm.migrateDatabase(); err != nil {
			return nil, err
		}

		instance = &dm
	}

	return instance, nil
}

func (db *DatabaseService) Login(username, password string) (bool, error) {
	var cnt int
	err := db.QueryRow(`select count(*) from User where Name = ? and Password = ? LIMIT 1`, username, hash(password)).Scan(&cnt)
	if err != nil {
		return false, errors.New("database problem")
	}
	return cnt != 0, nil
}

func (db *DatabaseService) CheckUsernameExists(username string) (bool, error) {
	var cnt int
	err := db.QueryRow(`select count(*) from User where Name = ? LIMIT 1`, username).Scan(&cnt)
	if err != nil {
		return false, errors.New("database problem: " + err.Error())
	}
	return cnt != 0, nil
}

func (db *DatabaseService) GetAll() ([]string, error) {
	var name, password string
	var res []string
	row, err := db.Query(`select Name, Password from User`)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		err := row.Scan(&name, &password)
		if err != nil {
			return nil, err
		}
		res = append(res, name+","+password+";")
	}
	return res, nil
}

func (db *DatabaseService) AddUser(username, email, password string) error {
	_, err := db.Exec(`INSERT INTO User (Name, Email, PASSWORD) VALUES (?, ?, ?)`, username, email, hash(password))
	return err
}

func (db *DatabaseService) Close() {
	db.DB.Close()
}

func (db *DatabaseService) migrateDatabase() error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS User(
    	Id integer PRIMARY KEY,
		Name varchar(255) UNIQUE NOT NULL,
		Email varchar(255),
		Password varchar(255) NOT NULL
    )`)
	if err != nil {
		return err
	}
	return nil
}

func hash(password string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}
