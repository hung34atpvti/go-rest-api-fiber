package storage

import "database/sql"

type Type int

const (
	POSTGRES = iota
	MYSQL
)

type Storage interface {
	Query (string, []any) (*sql.Rows, error)
}

var DB Storage

func NewStorage(t Type) error {
	if t == POSTGRES {
		DB = NewPostgresStorage()
	}
	if t == MYSQL {
		DB = NewMysqlStorage()
	}
	return nil
}
