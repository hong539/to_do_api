package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Task represents a to do item
type Task struct {
	ID      string `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

// tasks array represents in-memory storage for tasks
var tasks []Task

// GetTasksEndpoint returns all tasks
func GetTasksEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

// GetTaskEndpoint returns a single task by ID
func GetTaskEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range tasks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Task{})
}

// CreateTaskEndpoint creates a new task
func CreateTaskEndpoint(w http.ResponseWriter, req *http.Request) {
	var task Task
	_ = json.NewDecoder(req.Body).Decode(&task)
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

// DeleteTaskEndpoint deletes a task by ID
func DeleteTaskEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	// Initialize router
	router := mux.NewRouter()

	// Add sample data to the tasks array
	tasks = append(tasks, Task{ID: "1", Title: "Task 1", Content: "Content 1"})
	tasks = append(tasks, Task{ID: "2", Title: "Task 2", Content: "Content 2"})

	// Register endpoints
	router.HandleFunc("/tasks", GetTasksEndpoint).Methods("GET")
	router.HandleFunc("/tasks/{id}", GetTaskEndpoint).Methods("GET")
	router.HandleFunc("/tasks", CreateTaskEndpoint).Methods("POST")
	router.HandleFunc("/tasks/{id}", DeleteTaskEndpoint).Methods("DELETE")

	// Start the server
	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", router))
}