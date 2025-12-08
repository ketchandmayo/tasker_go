package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tasker_go/internal/task"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := []task.Task{
		{
			ID:          1,
			Title:       "Sample Task 1",
			Description: "This is the first sample task",
			Completed:   false,
		},
		{
			ID:          2,
			Title:       "Sample Task 2",
			Description: "This is the second sample task",
			Completed:   true,
		},
		{
			ID:          3,
			Title:       "Sample Task 3",
			Description: "This is the third sample task",
			Completed:   false,
		},
	}

	respondWithJSON(w, http.StatusOK, tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id == 10 {
		respondWithError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	t := task.Task{
		ID:          id,
		Title:       "Sample Task " + strconv.Itoa(id),
		Description: "This is the first sample task",
		Completed:   false,
	}

	respondWithJSON(w, http.StatusOK, t)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
