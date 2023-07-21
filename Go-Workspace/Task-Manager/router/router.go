package router

import (
	"github.com/gorilla/mux"
	"github.com/mdNoman21/Go/Beginner-Projects/Task-Manager/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/addTask", middleware.CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks", middleware.GetAllTask).Methods("GET")
	router.HandleFunc("/api/task/{id}", middleware.UpdateTask).Methods("PUT")
	router.HandleFunc("/api/task/{id}", middleware.DeleteTask).Methods("DELETE")
	router.HandleFunc("/api/tasks/completed", middleware.CompletedTasks).Methods("GET")
	router.HandleFunc("/api/task/{id}", middleware.SearchTask).Methods("GET")
	router.HandleFunc("/api/tasks/stat", middleware.TaskStats).Methods("GET")

	return router
}
