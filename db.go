package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
connStr := "host=" + os.Getenv("DB_HOST") +
	" port=" + os.Getenv("DB_PORT") +
	" user=" + os.Getenv("DB_USER") +
	" password=" + os.Getenv("DB_PASS") +
	" dbname=" + os.Getenv("DB_NAME") +
	" sslmode=require"
	
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	// ✅ Verify connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// ✅ NOW safe to configure
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	log.Println("Connected to RDS ✅")
}
