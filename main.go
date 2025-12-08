package main

import (
	"log"
	"net/http"
	"tasker_go/internal/router"
)

func main() {

	log.Println("Server started on http://localhost:8080/tasks")
	log.Fatal(http.ListenAndServe(":8080", router.NewRouter()))
}
