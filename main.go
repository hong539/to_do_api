package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/todos", GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", GetTodo).Methods("GET")
	router.HandleFunc("/todos", AddTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")

	log.Println("API server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
	log.Println("GetTodos: successfully get all todos.")
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, todo := range todos {
		if todo.ID == params["id"] {
			json.NewEncoder(w).Encode(todo)
			log.Println("GetTodo: successfully get todo with ID ", todo.ID)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})
	log.Println("GetTodo: todo with ID ", params["id"], " is not found.")
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
	log.Println("AddTodo: successfully add todo with title ", todo.Title)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, todo := range todos {
		if todo.ID == params["id"] {
			todos[index].Completed = true
			json.NewEncoder(w).Encode(todos[index])
			log.Println("UpdateTodo: successfully update todo with ID ", todo.ID)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})
	log.Println("UpdateTodo: todo with ID ", params["id"], " is not found.")
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, todo := range todos {
		if todo.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			json.NewEncoder(w).Encode(todo)
			log.Println("DeleteTodo: successfully delete todo with ID ", todo.ID)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})
	log.Println("DeleteTodo: todo with ID ", params["id"], " is not found.")
}