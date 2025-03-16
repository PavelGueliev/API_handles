package userService

import (
	"1/internal/taskService" // Исправлен импорт Task
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string             `json:"email"`    // Наш сервер будет ожидать json c полем text
	Password string             `json:"password"` // В GO используем CamelCase, в Json - snake
	Tasks    []taskService.Task `json:"tasks" gorm:"foreignKey:UserID"`
}
