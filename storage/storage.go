package storage

import "database/sql"

type Storage interface {
	Query(query string) (*sql.Rows, error)
}

var DB Storage

func New(dbType string) error {
	if dbType == "postgres" {
		DB = NewPostgresStorage()
	}
	if dbType == "mysql" {
		DB = NewMysqlStorage()
	}
	return nil
}
