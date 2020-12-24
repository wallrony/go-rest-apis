package database

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	db *sql.DB
)

// InitializeDB function iniatalize DB instance
// connecting to database with the defined vars.
func InitializeDB() {
	var err error

	db, err = sql.Open("postgres", connString())

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
}

func connString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("user=%s dbname=%s port=%s password=%s host=%s sslmode=disable", user, dbname, port, pass, host)
	return connStr
}

// GetDB function returns database instance.
func GetDB() *sql.DB {
	return db
}
