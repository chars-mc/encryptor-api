package database

import (
	"database/sql"

	"github.com/chars-mc/encryptor-api/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLClient struct {
	*sql.DB
}

func NewMySQLClient(config *config.DBConfig) (*MySQLClient, error) {
	db, err := sql.Open("mysql", config.GetDSN())
	if err != nil {
		return nil, err
	}
	return &MySQLClient{db}, nil
}
