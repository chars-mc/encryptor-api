package config

import "fmt"

type DBConfig struct {
	host     string
	port     string
	username string
	password string
	database string
}

func NewDBConfig(host, port, username, password, database string) *DBConfig {
	return &DBConfig{
		host:     host,
		port:     port,
		username: username,
		password: password,
		database: database,
	}
}

// GetDSN returns the Data Source Name to connect to database server
// username:password@protocol(address)/dbname?parseTime=true
func (db *DBConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		db.username,
		db.password,
		db.host,
		db.port,
		db.database,
	)
}
