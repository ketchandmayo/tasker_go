package main

import (
	"log"
	"net/http"
	"tasker_go/internal/config"
	"tasker_go/internal/models"
	"tasker_go/internal/transport/http/router"
)

func main() {
	config.ConnectDB()
	err := config.DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Migration error")
	}

	log.Println("Server started on http://localhost:8080/tasks")
	log.Fatal(http.ListenAndServe(":8080", router.MainRouter()))
}
