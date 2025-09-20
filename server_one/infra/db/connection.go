package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	return "user=minar password=1313 dbname=serverone host=localhost port=5432 sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	connStr := GetConnectionString()

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
		return nil, err
	}

	return db, nil
}
