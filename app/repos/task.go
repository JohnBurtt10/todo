package repos

import (
	"errors"

	"github.com/JohnBurtt10/go/app/models"
	"github.com/JohnBurtt10/go/database"
	"gorm.io/gorm"
)

func GetTasks(userID uint) (*[]models.Task, error) {
	// var tasks []models.Task
	var user models.User
	database.DBConn.First(&user, userID)
	var tasks []models.Task
	err := database.DBConn.Model(&user).Association("Tasks").Find(&tasks)
	return &tasks, err
}

func DeleteTask(taskID uint) error {
	var item models.Task
	// Delete the note and return error if encountered
	err := database.DBConn.Delete(&item, taskID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("task not found")
	}
	return nil
}

func CreateTask(TaskName string, Assignee string, IsDone bool, UserID uint) (*models.Task, error) {
	var task = &models.Task{TaskName: TaskName, Assignee: Assignee, IsDone: IsDone, UserID: UserID}
	if err := database.DBConn.Create(&task).Error; err != nil {
		return nil, err
	}
	//
	return task, nil
	// ...
}
