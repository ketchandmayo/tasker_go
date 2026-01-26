package app

import (
	"log"
	"net/http"
	"tasker_go/internal/config"
	"tasker_go/internal/repository/gorm"
	"tasker_go/internal/service"
	"tasker_go/internal/transport/http/handler"
	"tasker_go/internal/transport/http/router"

	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",
			"http://localhost:63342",
		},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Authorization",
			"Content-Type",
		},
	})

	err := http.ListenAndServe(":8080", c.Handler(r))
	if err != nil {
		log.Printf("Server error")
	}

}
