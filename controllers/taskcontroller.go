package controllers

import (
	"go-crud-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	models.DB.Preload("Category").Preload("Task").Find(&tasks)
	c.JSON(200, tasks)
	return
}

func CreateTask(c *gin.Context) {
	var task models.Task
	validate := validator.New()
	err := validate.Struct(task)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, errors)
		return
	}
	err = c.ShouldBindJSON(&task)
	res := models.DB.Create(&task)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(201, task)
	return
}

func GetTaskById(c *gin.Context) {
	taskId := c.Param("taskId")
	var task models.Task
	res := models.DB.Preload("Category").Find(&task, taskId)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, &task)
	return
}

func DeleteTask(c *gin.Context) {
	taskId := c.Param("taskId")
	var task models.Task
	res := models.DB.Find(&task, taskId)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	models.DB.Delete(&task, taskId)
	c.JSON(http.StatusNoContent, task)
	return
}

func UpdateTask(c *gin.Context) {
	taskId := c.Param("taskId")
	var task models.Task
	res := models.DB.First(&task, taskId)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.ShouldBindJSON(&task)
	models.DB.Save(&task)
	c.JSON(http.StatusOK, &task)
}
