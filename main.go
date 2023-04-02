package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

var db *gorm.DB

func main() {
	// 連接資料庫
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 自動建立資料表
	db.AutoMigrate(&Todo{})

	// 設定router
	router := mux.NewRouter()

	// 設定API endpoint
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	router.HandleFunc("/todos", addTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	// 啟動API服務
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// 取得所有待辦事項列表
func getTodos(w http.ResponseWriter, r *http.Request) {
	var todos []Todo
	db.Find(&todos)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

// 取得特定id的待辦事項
func getTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo Todo
	db.First(&todo, params["id"])

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}	