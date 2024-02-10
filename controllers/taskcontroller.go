package controllers

import (
	"encoding/json"
	"fmt"
	"go-crud-api/models"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	models.DB.Preload("Category").Preload("Task").Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	validate := validator.New()
	json.NewDecoder(r.Body).Decode(&task)
	err := validate.Struct(task)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return
	}
	res := models.DB.Create(&task)
	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["taskId"]
	var task models.Task
	res := models.DB.Preload("Category").Find(&task, taskId)
	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["taskId"]
	var task models.Task
	res := models.DB.Find(&task, taskId)
	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	models.DB.Delete(&task, taskId)
	w.WriteHeader(http.StatusNoContent)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["taskId"]
	var task models.Task
	res := models.DB.First(&task, taskId)
	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&task)
	models.DB.Save(&task)
	json.NewEncoder(w).Encode(task)
}
