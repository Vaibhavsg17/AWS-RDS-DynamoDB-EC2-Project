package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	log.Println("Connected to RDS ✅")
}

func createTables() {

	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT,
		email TEXT
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS projects (
		id TEXT PRIMARY KEY,
		name TEXT,
		user_id TEXT
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tables ensured ✅")
}
