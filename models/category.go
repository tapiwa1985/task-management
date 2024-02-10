package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID           uint   `gorm:"primary key"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}
