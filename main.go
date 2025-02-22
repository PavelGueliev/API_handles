package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type requestBody struct {
	Message string `json:"message"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var req requestBody
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	task = req.Message
	fmt.Fprintf(w, "Task updated to: %s", task)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if task == "" {
		fmt.Fprintln(w, "Hello, World!")
	} else {
		fmt.Fprintf(w, "Hello, %s", task)
	}
}

func main() {
	router := mux.NewRouter()
	// наше приложение будет слушать запросы на localhost:8080/api/hello
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")

	// Обработка POST запроса
	router.HandleFunc("/api/task", PostHandler).Methods("POST")

	// Запуск сервера
	http.ListenAndServe(":8080", router)
}
