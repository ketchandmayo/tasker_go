package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"tasker_go/internal/service"
	"tasker_go/internal/transport/http/dto"
	"tasker_go/internal/transport/http/middleware"

	"github.com/gorilla/mux"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	userID := uint(1) // позже достанем из ctx

	var req dto.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if err := dto.Validate.Struct(req); err != nil {
		respondWithError(w, http.StatusBadRequest, dto.FormatValidationError(err))
		return
	}

	task := req.ToModel(userID)

	id, err := h.taskService.Create(r.Context(), userID, &task)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal error")
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]uint{"id": id})
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	userID := uint(1)

	tasks, err := h.taskService.List(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal error")
		return
	}

	resp := make([]dto.TaskListItemResponse, 0, len(tasks))
	for _, t := range tasks {
		resp = append(resp, dto.TaskToListItem(t))
	}

	respondWithJSON(w, http.StatusOK, resp)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	userID := uint(1)

	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid id")
		return
	}

	task, err := h.taskService.Get(r.Context(), userID, uint(id))
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			respondWithError(w, http.StatusNotFound, "task not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "internal error")
		return
	}

	resp := dto.TaskToDetailResponse(*task)
	respondWithJSON(w, http.StatusOK, resp)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	userID := uint(1)

	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req dto.PatchTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if err := dto.Validate.Struct(req); err != nil {
		respondWithError(w, http.StatusBadRequest, dto.FormatValidationError(err))
		return
	}

	task, err := h.taskService.Update(r.Context(), userID, uint(id), &req)
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			respondWithError(w, http.StatusNotFound, "task not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "internal error")
		return
	}

	resp := dto.TaskToDetailResponse(*task)
	respondWithJSON(w, http.StatusOK, resp)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	userID := uint(1)

	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.taskService.Delete(r.Context(), userID, uint(id))
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			respondWithError(w, http.StatusNotFound, "task not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "internal error")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func UserIDFromContext(ctx context.Context) uint {
	id, ok := ctx.Value(middleware.UserIDKey).(uint)
	if !ok {
		panic("userID not found in context")
	}
	return id
}
