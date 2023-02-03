package storage

import (
	"database/sql"
	"rest-api/database/mysql"
	"rest-api/database/postgres"
)

type Storage interface {
	Query(query string) error
}

var DB Storage

func New(dbType string) error {
	if dbType == "postgres" {
		DB = NewPostgresStorage()
	}
	if dbType == "mysql" {
		DB = mysql.DB
	}
	return nil
}
