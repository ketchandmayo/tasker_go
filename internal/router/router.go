package router

import (
	"tasker_go/internal/handler"

	"github.com/gorilla/mux"
)

func MainRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", handler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id:[0-9]+}", handler.GetTask).Methods("GET")
	r.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id:[0-9]+}", handler.UpdateTask).Methods("PATCH")
	//r.HandleFunc("/tasks/{id:[0-9]+}", handler.DeleteTask).Methods("DELETE")

	return r
}
