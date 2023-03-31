package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

type Task struct {
    ID        string `json:"id,omitempty"`
    Name      string `json:"name,omitempty"`
    Completed bool   `json:"completed,omitempty"`
}

var tasks []Task

func GetTasks(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task Task
    _ = json.NewDecoder(r.Body).Decode(&task)
    tasks = append(tasks, task)
    json.NewEncoder(w).Encode(task)
}

func main() {
    router := mux.NewRouter()

    tasks = append(tasks, Task{ID: "1", Name: "Task 1", Completed: false})
    tasks = append(tasks, Task{ID: "2", Name: "Task 2", Completed: true})

    router.HandleFunc("/tasks", GetTasks).Methods("GET")
    router.HandleFunc("/tasks", CreateTask).Methods("POST")

    log.Fatal(http.ListenAndServe(":8000", router))
}