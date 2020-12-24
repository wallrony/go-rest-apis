package db

import (
	"database/sql"
	"os"

	f "fmt"
)

var (
	db *sql.DB
)

// InitializeDB function initialize postgres db instance.
func InitializeDB() {
	var err error

	db, err = sql.Open("postgres", connStr())

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	f.Printf("Successfully connected to database!\n")
}

func connStr() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	connStr := f.Sprintf("user=%s dbname=%s port=%s password=%s host=%s sslmode=disable", user, dbname, port, password, host)

	return connStr
}

// GetDB function returns the database instance.
func GetDB() *sql.DB {
	return db
}
