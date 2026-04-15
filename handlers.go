package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// ---------------- USERS ----------------

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user.ID = uuid.New().String()

	query := `INSERT INTO users (id, name, email) VALUES ($1, $2, $3)`
	_, err := DB.Exec(query, user.ID, user.Name, user.Email)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, err := DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("User deleted"))
}

// ---------------- PROJECTS ----------------

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project Project

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	project.ID = uuid.New().String()

	query := `INSERT INTO projects (id, name, user_id) VALUES ($1, $2, $3)`
	_, err := DB.Exec(query, project.ID, project.Name, project.UserID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(project)
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, name, user_id FROM projects")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var p Project
		rows.Scan(&p.ID, &p.Name, &p.UserID)
		projects = append(projects, p)
	}

	json.NewEncoder(w).Encode(projects)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, err := DB.Exec("DELETE FROM projects WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("Project deleted"))
}

// ---------------- TASKS (DynamoDB) ----------------
var TableName = "Tasks"

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task

	// Decode JSON body
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate UUID for primary key
	task.TaskID = uuid.New().String()

	// Marshal to DynamoDB attribute values
	item, err := attributevalue.MarshalMap(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Put item in DynamoDB
	_, err = DynamoClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: &TableName,
		Item:      item,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created task as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	result, err := DynamoClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: &TableName,
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var tasks []Task
	attributevalue.UnmarshalListOfMaps(result.Items, &tasks)

	json.NewEncoder(w).Encode(tasks)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, err := DynamoClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: &TableName,
		Key: map[string]types.AttributeValue{
			"task_id": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("Task deleted"))
}
