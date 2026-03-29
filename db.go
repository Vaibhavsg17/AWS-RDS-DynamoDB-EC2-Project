package main

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
 	"os"
)

var DB *sql.DB

func InitDB() {
	connStr := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASS") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	log.Println("Connected to RDS ✅")
}