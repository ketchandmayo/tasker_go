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
	repo := gorm.NewTaskRepository(db)
	svc := service.NewTaskService(repo)
	h := handler.New(svc)
	r := router.New(h)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Printf("Server error")
	}
}
