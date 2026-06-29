package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	// connStr := "host=localhost port=5432 user=postgres password=109798 dbname=postgres sslmode=disable"
	connStr := os.Getenv("ConnStr")
	if connStr == "" {
		log.Fatal("Db creds missing")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	fmt.Println("Connected to PostgreSQL")

	return db
}
