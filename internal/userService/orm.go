package userService

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`    // Наш сервер будет ожидать json c полем text
	Password string `json:"password"` // В GO используем CamelCase, в Json - snake
}
