package main

import (
	"1/internal/database"
	"1/internal/handlers"
	"1/internal/taskService"
	"1/internal/userService"
	"1/internal/web/tasks"
	"1/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()

	taskRepo := taskService.NewTaskRepository(database.DB)
	userRepo := userService.NewUserRepository(database.DB)

	taskService := taskService.NewService(taskRepo)
	userService := userService.NewUserService(userRepo)

	taskHandlers := handlers.NewTaskHandler(taskService, userService)
	userHandlers := handlers.NewUserHandler(userService)
	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictTaskHandler := tasks.NewStrictHandler(taskHandlers, nil) // Правильный handler для задач
	strictUserHandler := users.NewStrictHandler(userHandlers, nil) // Новый handler для пользователей

	tasks.RegisterHandlers(e, strictTaskHandler) // Регистрируем `tasks`
	users.RegisterHandlers(e, strictUserHandler) // Регистрируем `users`

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
