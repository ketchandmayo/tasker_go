package router

import (
	"net/http"
	"tasker_go/internal/transport/http/handler"
	"tasker_go/internal/transport/http/middleware"

	"github.com/gorilla/mux"
)

func New(h *handler.Handler) http.Handler {
	r := mux.NewRouter()

	// r.HandleFunc("/login", h.Login).Methods(http.MethodPost)
	// r.HandleFunc("/register", h.Register).Methods(http.MethodPost)

	auth := r.PathPrefix("/").Subrouter()
	auth.Use(middleware.Auth)

	auth.HandleFunc("/tasks", h.CreateTask).Methods(http.MethodPost)
	auth.HandleFunc("/tasks", h.GetTasks).Methods(http.MethodGet)
	auth.HandleFunc("/tasks/{id}", h.GetTask).Methods(http.MethodGet)
	auth.HandleFunc("/tasks/{id}", h.UpdateTask).Methods(http.MethodPatch)
	auth.HandleFunc("/tasks/{id}", h.DeleteTask).Methods(http.MethodDelete)

	return r
}
