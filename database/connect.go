package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

func connectDB() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
		return nil, err
	}

	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
		return nil, err
	}

	return db, nil
}
