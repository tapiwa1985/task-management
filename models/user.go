package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint   `gorm:"primary key"`
	Firstname    string `json:"first_name" validate:"required"`
	Lastname     string `json:"last_name" validate:"required"`
	Email        string `json:"email" validate:"required"`
	Password     string `json:"password" validate:"required"`
	MobileNumber string `json:"mobile_number" validate:"required"`
}

func (user *User) HashPassword(password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(passwordHash)
	return nil
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
