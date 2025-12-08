package main

import (
	"log"
	"net/http"
	"tasker_go/internal/handler"
)

func main() {
	http.HandleFunc("/tasks", handler.TaskHandler)

	log.Println("Server started on http://localhost:8080/tasks")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
