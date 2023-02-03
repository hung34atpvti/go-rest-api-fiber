package storage

import (
	"database/sql"
	"rest-api/database/postgres"
)

var db = postgres.DB
type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	var err error

	stg := new(PostgresStorage)

	stg.db, err = db, nil

	return stg, err
}

func Query(query string)  error {
	db.Query(query)
	return nil
}