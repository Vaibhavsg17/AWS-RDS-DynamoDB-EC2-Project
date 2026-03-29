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
	InitDynamo()

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