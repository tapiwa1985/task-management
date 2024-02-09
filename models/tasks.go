package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID              uint      `gorm:"primary key"`
	TaskName        string    `json:"task_name" validate:"required"`
	TaskDescription string    `json:"description" validate:"required"`
	AssignedUser    string    `json:"assigned_user" validate:"required"`
	Deadline        time.Time `json:"deadline"`
}
