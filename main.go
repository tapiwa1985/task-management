package main

import (
	"go-crud-api/controllers"
	"go-crud-api/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func main() {
	models.ConnectToDatabase()
	router = mux.NewRouter()
	RegisterRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(router)))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application-json")
		next.ServeHTTP(w, r)
	})
}

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{taskId}", controllers.GetTaskById).Methods("GET")
	r.HandleFunc("/tasks/{taskId}", controllers.DeleteTask).Methods("DELETE")
}
