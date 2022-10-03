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
	database.DBConn.First(&user, 1)
	var tasks []models.Task

	// err := database.DBConn.Model(&models.User{}).Preload("Tasks").Find(&tasks).Error
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return nil, errors.New("task not found")
	// }
	// for _, m := range tasks {
	// 	fmt.Println("Task name:", m.TaskName, "Assignee:", m.Assignee, "IsDone:", m.IsDone, "UserID:", m.UserID)
	// }
	err := database.DBConn.Model(&user).Association("Tasks").Find(&tasks)
	return &tasks, err
}

func DeleteTask(taskId uint) error {
	var item models.Task
	// Delete the note and return error if encountered
	err := database.DBConn.Delete(&item, taskId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("task not found")
	}
	return nil
}

func CreateTask(TaskName string, Assignee string, IsDone bool, UserID uint) error {
	var item = models.Task{TaskName: TaskName, Assignee: Assignee, IsDone: IsDone, UserID: UserID}
	database.DBConn.Select("Title", "UserID").Create(&item)
	return nil
	// ...
}
