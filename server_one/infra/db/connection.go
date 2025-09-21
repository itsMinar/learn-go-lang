package db

import (
	"fmt"
	"my_server/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DBConfig) string {
	sslMode := "disable"
	if cnf.EnableSSLMode {
		sslMode = "enable"
	}
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cnf.User, cnf.Password, cnf.DBName, cnf.Host, cnf.Port, sslMode)
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	connStr := GetConnectionString(cnf)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
		return nil, err
	}

	return db, nil
}
