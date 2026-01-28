package router

import (
	"net/http"
	"tasker_go/internal/transport/http/handler"
	"tasker_go/internal/transport/http/middleware"

	"github.com/gorilla/mux"
)

func New(taskHandler *handler.TaskHandler, authHandler *handler.AuthHandler, judgeHandler *handler.JudgeHandler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
	r.HandleFunc("/register", authHandler.Register).Methods(http.MethodPost)

	auth := r.PathPrefix("/").Subrouter()
	auth.Use(middleware.Auth)

	auth.HandleFunc("/tasks", taskHandler.CreateTask).Methods(http.MethodPost)
	auth.HandleFunc("/tasks", taskHandler.GetTasks).Methods(http.MethodGet)
	auth.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods(http.MethodGet)
	auth.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods(http.MethodPatch)
	auth.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods(http.MethodDelete)
	auth.HandleFunc("/tasks/{id}/judge", judgeHandler.GetJudge).Methods(http.MethodGet)

	return r
}
