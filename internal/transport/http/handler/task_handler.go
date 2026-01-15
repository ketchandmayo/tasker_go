package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"tasker_go/internal/config"
	"tasker_go/internal/models"
	"tasker_go/internal/transport/http/dto"

	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err.Error())
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := dto.Validate.Struct(req); err != nil {
		log.Printf("Validation error: %v", err)
		respondWithError(w, http.StatusBadRequest, dto.FormatValidationError(err))
		return
	}

	task := req.ToModel(uint(1))

	if err := config.DB.Create(&task).Error; err != nil {
		log.Printf("DB error: %v", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Received task: %+v", task)
	respondWithJSON(w, http.StatusOK, map[string]uint{"id": task.ID})
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := make([]models.Task, 0)
	userId := uint(1)

	if err := config.DB.
		Where("user_id = ?", userId).
		Order("created_at desc").
		Find(&tasks).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var responses []dto.TaskListItemResponse
	for _, t := range tasks {
		responses = append(responses, dto.TaskToListItem(t))
	}

	respondWithJSON(w, http.StatusOK, responses)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	response := dto.TaskToDetailResponse(task)

	log.Printf("Requested task: %+v", task)
	respondWithJSON(w, http.StatusOK, response)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userID := uint(1)

	var req dto.PatchTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err.Error())
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var task models.Task
	if err := config.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&task).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "task not found")
		return
	}

	if err := dto.Validate.Struct(req); err != nil {
		log.Printf("Validation error: %v", err)
		respondWithError(w, http.StatusBadRequest, dto.FormatValidationError(err))
		return
	}
	req.Apply(&task)
	if err := config.DB.Save(&task).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := dto.TaskToDetailResponse(task)
	respondWithJSON(w, http.StatusOK, response)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userID := uint(1)

	var task models.Task
	if err := config.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&task).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "task not found")
		return
	}

	if err := config.DB.Delete(&task).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}
