package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Load env (works locally, ignored in Docker if env already set)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found (using system env)")
	}

	// Init PostgreSQL (RDS)
	InitDB()

	// Create tables safely (idempotent)
	createTables()

	// Init DynamoDB
	InitDynamo()
	TestDynamoConnection()


	http.HandleFunc("/health", HealthCheck)
	http.HandleFunc("/users", CreateUser)
	http.HandleFunc("/get-users", GetUsers)
	http.HandleFunc("/delete-user", DeleteUser)

	http.HandleFunc("/projects", CreateProject)
	http.HandleFunc("/get-projects", GetProjects)
	http.HandleFunc("/delete-project", DeleteProject)

	http.HandleFunc("/tasks", CreateTask)
	http.HandleFunc("/get-tasks", GetTasks)
	http.HandleFunc("/delete-task", DeleteTask)

port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on :" + port + " 🚀")

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
