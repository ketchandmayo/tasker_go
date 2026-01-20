package app

import (
	"log"
	"net/http"
	"tasker_go/internal/config"
	"tasker_go/internal/repository/gorm"
	"tasker_go/internal/service"
	"tasker_go/internal/transport/http/handler"
	"tasker_go/internal/transport/http/router"
)

func Run() {
	db := config.ConnectDB()
	taskRepo := gorm.NewTaskRepository(db)
	userRepo := gorm.NewUserRepository(db)

	taskService := service.NewTaskService(taskRepo)
	authService := service.NewAuthService(userRepo)

	taskHandler := handler.NewTaskHandler(taskService)
	authHandler := handler.NewAuthHandler(authService)

	r := router.New(taskHandler, authHandler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Printf("Server error")
	}
}
