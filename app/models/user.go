package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Username  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Tasks     []Task `json:"tasks"`
}

type UserDTO struct {
	FirstName       *string `json:"firstName"`
	LastName        *string `json:"lastName"`
	Username        string  `gorm:"not null"`
	CurrentPassword *string `json:"currentPassword"`
	Password        *string `json:"password" validate:"min:8"`
}

// TODO: dto and validate in controller

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Tasks    []Task `json:"tasks"`
}
