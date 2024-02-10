package controllers

import (
	"go-crud-api/auth"
	"go-crud-api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var u models.User
	c.ShouldBindJSON(&u)
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, errors)
		return
	}
	err = u.HashPassword(u.Password)
	if err != nil {
		log.Fatal(err)
		return
	}
	res := models.DB.Create(&u)
	if res.RowsAffected != 0 {
		c.JSON(http.StatusCreated, &u)
		return
	} else {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
}

func Login(c *gin.Context) {
	var request LoginRequest
	var user models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	record := models.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	loginError := user.CheckPassword(request.Password)
	if loginError != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	token, err := auth.GenerateJWT(request.Email)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(200, gin.H{"token": token})
}
