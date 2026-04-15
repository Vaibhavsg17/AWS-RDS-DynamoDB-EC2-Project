package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	InitDB()

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

	log.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
