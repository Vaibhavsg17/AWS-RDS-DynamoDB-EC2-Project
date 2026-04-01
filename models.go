package main

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Project struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

type Task struct {
	TaskID    string `json:"task_id" dynamodbav:"task_id"`   // PRIMARY KEY
	ProjectID string `json:"project_id" dynamodbav:"project_id"`
	Title     string `json:"title" dynamodbav:"title"`
	Status    string `json:"status" dynamodbav:"status"`
}