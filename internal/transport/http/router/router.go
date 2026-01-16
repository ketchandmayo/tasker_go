package router

import (
	"net/http"
	"tasker_go/internal/transport/http/handler"

	"github.com/gorilla/mux"
)

func New(h *handler.Handler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", h.CreateTask).Methods(http.MethodPost)
	r.HandleFunc("/tasks", h.GetTasks).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id}", h.GetTask).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id}", h.UpdateTask).Methods(http.MethodPatch)
	r.HandleFunc("/tasks/{id}", h.DeleteTask).Methods(http.MethodDelete)

	return r
}
