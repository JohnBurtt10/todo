package models

import (
	"gorm.io/gorm"
)

// Task struct defines the task
type Task struct {
	gorm.Model
	Title    string `json:"title"`
	Assignee string `json:"assignee"`
	IsDone   bool   `json:"isDone"`
	UserID   uint   `json:"userID"`
}

type TaskDTO struct {
	ID       *uint  `json:"id"`
	Title    string `json:"title"`
	Assignee string `json:"assignee"`
	IsDone   bool   `json:"isDone"`
	UserID   uint   `json:"userID"`
}

// TODO: add dto and use BodyParseAndValidate in controller

type TaskResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Assignee string `json:"assignee"`
	IsDone   bool   `json:"isDone"`
	UserID   uint   `json:"userID"`
}
