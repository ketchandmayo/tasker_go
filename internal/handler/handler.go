package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"tasker_go/internal/task"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	t := task.Task{
		ID:          1,
		Title:       "Sample Task",
		Description: "This is a sample task",
		Completed:   false,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(t)

	if err != nil {
		log.Fatal(err)
	}
}
