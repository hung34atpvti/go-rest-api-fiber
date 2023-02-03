package storage

import (
	"database/sql"
	"rest-api/database/postgres"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() *PostgresStorage {
	s := new(PostgresStorage)
	s.db = postgres.DB
	return s
}

func (s *PostgresStorage) Query(query string)  (*sql.Rows, error) {
	return s.db.Query(query)
}