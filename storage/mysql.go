package storage

import (
	"database/sql"
	"rest-api/database/mysql"
)

type MysqlStorage struct {}

func NewMysqlStorage() *MysqlStorage {
	return &MysqlStorage{}
}

func (s *MysqlStorage) Query(query string)  (*sql.Rows, error) {
	return mysql.DB.Query(query)
}