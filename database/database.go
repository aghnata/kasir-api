package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" //_ means dipake magicnya tp ga dipanggil d project
)

func InitDB(connectionString string) (*sql.DB, error) {
	//open connection to database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	//Test Connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	//set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
	return db, nil
}
