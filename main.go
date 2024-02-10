package main

import (
	"go-crud-api/controllers"
	"go-crud-api/middleware"
	"go-crud-api/models"
	"log"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	models.ConnectToDatabase()
	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.POST("/registration", controllers.Register)
		api.POST("/auth/token", controllers.Login)
		api.GET("/tasks", middleware.Auth(), controllers.GetTasks)
		api.GET("/tasks/:taskId", middleware.Auth(), controllers.GetTaskById)
		api.PUT("/tasks/:taskId", middleware.Auth(), controllers.UpdateTask)
		api.DELETE("/tasks:taskId", middleware.Auth(), controllers.DeleteTask)
		api.POST("/tasks", middleware.Auth(), controllers.CreateTask)
	}
	r.Run(":8000")
}
