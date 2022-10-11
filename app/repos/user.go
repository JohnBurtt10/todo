package repos

import (
	"errors"

	"github.com/JohnBurtt10/go/app/models"
	"github.com/JohnBurtt10/go/database"
	"gorm.io/gorm"
)

func GetUserByUsername(username string) (*models.User, error) {
	user := models.User{}
	err := database.DBConn.Model(models.User{Username: username}).First(&user).Error
	return &user, err
}

func GetUserByID(userID uint) (*models.User, error) {
	user := &models.User{}
	if err := database.DBConn.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Creates an user in the user's table
func CreateUser(user *models.User) error {
	err := database.DBConn.Model(&models.User{}).Where("username = ?", user.Username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return database.DBConn.Create(user).Error
	} else {
		return errors.New("User with that username already exists")
	}
}

func ChangeUserPassword(user *models.User, newPassword string) error {
	err := database.DBConn.Model(&models.User{}).Where("Username = ?", user.Username).Update("Password", newPassword).Error
	if err != nil {
		return err
	}
	return nil
}
