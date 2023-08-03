package database

import (
	"NAS-Server-Web/operations"
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseManager struct {
	*sql.DB
}

var instance *DatabaseManager = nil

func GetDatabase(databaseLocation string) (*DatabaseManager, error) {
	if instance == nil {
		db, err := sql.Open("sqlite3", databaseLocation)
		if err != nil {
			return nil, err
		}

		dm := DatabaseManager{db}
		if err = dm.migrateDatabase(); err != nil {
			return nil, err
		}

		instance = &dm
	}

	return instance, nil
}

func (db *DatabaseManager) Login(username, password string) (bool, error) {
	var cnt int
	err := db.QueryRow(`select count(*) from User where Name = ? and Password = ? LIMIT 1`, username, operations.Hash(password)).Scan(&cnt)
	if err != nil {
		return false, errors.New("database problem")
	}
	return cnt != 0, nil
}

func (db *DatabaseManager) CheckUsernameExists(username string) (bool, error) {
	var cnt int
	err := db.QueryRow(`select count(*) from User where Name = ? LIMIT 1`, username).Scan(&cnt)
	if err != nil {
		return false, errors.New("database problem: " + err.Error())
	}
	return cnt != 0, nil
}

func (db *DatabaseManager) GetAll() ([]string, error) {
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

func (db *DatabaseManager) migrateDatabase() error {
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

func (db *DatabaseManager) Close() {
	db.Close()
}
