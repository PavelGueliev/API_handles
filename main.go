package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var task Task

	// Декодируем JSON из тела запроса
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Записываем задачу в БД
	result := DB.Create(&task)
	if result.Error != nil {
		http.Error(w, "Failed to save task", http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	var tasks []Task

	// Получаем все записи из БД
	result := DB.Find(&tasks)
	if result.Error != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	// Отправляем JSON с задачами
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func PatchMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Конвертируем ID из строки в int
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var task Task
	result := DB.First(&task, id) // Находим задачу по ID
	if result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	// Обновляем только переданные поля
	if updatedTask.Task != "" {
		task.Task = updatedTask.Task
	}
	task.IsDone = updatedTask.IsDone

	DB.Save(&task) // Сохраняем изменения в БД

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) // Конвертируем ID в int
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task Task
	result := DB.First(&task, id) // Ищем задачу
	if result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	DB.Delete(&task) // Удаляем задачу

	w.WriteHeader(http.StatusNoContent) // 204 No Content (без тела ответа)
}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	// наше приложение будет слушать запросы на localhost:8080/api/hello
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")

	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	router.HandleFunc("/api/messages/{id:[0-9]+}", PatchMessage).Methods("PATCH")
	router.HandleFunc("/api/messages/{id:[0-9]+}", DeleteMessages).Methods("DELETE")

	// Обработка POST запроса
	router.HandleFunc("/api/task", PostHandler).Methods("POST")

	// Запуск сервера
	http.ListenAndServe(":8080", router)
}
