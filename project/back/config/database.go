package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func SetupPostgres() (*sql.DB, error) {
	connStr := "postgres://postgres:ftn@localhost/bezbednost?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSQL database!")
	return db, nil
}
