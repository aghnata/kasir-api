package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// InitDB initializes and returns a database connection
func InitDB(connectionString string) (*sql.DB, error) {
	// Open connection to database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
	return db, nil
}
