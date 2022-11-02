package repos

import (
	"errors"

	"github.com/JohnBurtt10/go/app/models"
	"github.com/JohnBurtt10/go/database"
	"gorm.io/gorm"
)

func GetTasks(userID uint) (*[]models.Task, error) {
	var user models.User
	database.DBConn.First(&user, userID)
	var tasks []models.Task
	err := database.DBConn.Model(&user).Association("Tasks").Find(&tasks)
	return &tasks, err
}

func UpdateTask(taskID uint, title string, assignee string, isDone bool) (*models.Task, error) {
	var task = models.Task{Title: title, Assignee: assignee, IsDone: isDone}
	task.ID = taskID
	err := database.DBConn.Model(&task).Select("Title", "Assignee", "IsDone").Where("id = ?", taskID).Updates(models.Task{Title: title, Assignee: assignee, IsDone: isDone}).Error
	return &task, err

}

func DeleteTask(taskID uint) error {
	var item models.Task
	err := database.DBConn.Delete(&item, taskID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("task not found")
	}
	return nil
}

func CreateTask(title string, assignee string, isDone bool, userID uint) (*models.Task, error) {
	var task = &models.Task{Title: title, Assignee: assignee, IsDone: isDone, UserID: userID} // this will create a NULL value in the db if isDOne == false
	if err := database.DBConn.Create(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}
