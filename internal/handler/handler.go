package handler

import (
	"encoding/json"
	"log"
	"net/http"
	//"strconv"
	//"tasker_go/internal/task"
	//
	//"github.com/gorilla/mux"

	"tasker_go/internal/config"
	"tasker_go/internal/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		log.Printf("Error decoding request body: %v", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := config.DB.Create(&task).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Received task: %+v", task)
	respondWithJSON(w, http.StatusOK, task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	if err := config.DB.Find(&tasks).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {

	//respondWithJSON(w, http.StatusOK, t)
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
	_, err = w.Write(response)
	if err != nil {
		return
	}
}
