package storage

import (
	"database/sql"
	"rest-api/database/postgres"
)

type PostgresStorage struct {}

func NewPostgresStorage() *PostgresStorage {
	return &PostgresStorage{}
}

func (s *PostgresStorage) Query(query string, params []any)  (*sql.Rows, error) {
	return postgres.DB.Query(query, params...)
}