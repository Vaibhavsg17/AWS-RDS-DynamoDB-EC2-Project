package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found (using system env)")
	}

	InitDB()

	_, err = DB.Exec(`
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

	http.HandleFunc("/users", CreateUser)
	http.HandleFunc("/get-users", GetUsers)
	http.HandleFunc("/delete-user", DeleteUser)

	http.HandleFunc("/projects", CreateProject)
	http.HandleFunc("/get-projects", GetProjects)
	http.HandleFunc("/delete-project", DeleteProject)

	http.HandleFunc("/tasks", CreateTask)
	http.HandleFunc("/get-tasks", GetTasks)
	http.HandleFunc("/delete-task", DeleteTask)

	log.Println("Server running on :8080 🚀")
	http.ListenAndServe(":8080", nil)
}
