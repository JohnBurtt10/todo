package repos

import (
	"errors"

	"github.com/JohnBurtt10/go/app/models"
	"github.com/JohnBurtt10/go/app/utils/password/argon2id"
	"github.com/JohnBurtt10/go/database"
	"gorm.io/gorm"
)

// FindUser searches the user's table with the condition given
func FindUser(dest interface{}, conds ...interface{}) error {
	// works
	return database.DBConn.Model(&models.User{}).Take(dest, conds...).Error
}

// FindUserByEmail searches the user's table with the email given
func FindUserByUsername(dest interface{}, username string) error {
	return FindUser(dest, "username = ?", username)
}

func FindUserByID(dest interface{}, ID uint) error {
	return FindUser(dest, "ID = ?", ID)
}

// Creates an user in the user's table
func CreateUser(user *models.User) error {
	err := FindUserByUsername(nil, user.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return database.DBConn.Create(user).Error
	} else {
		return errors.New("User with that username already exists")
	}
}

func ChangeUserPassword(user *models.User, newPassword string) error {
	argon2ID := argon2id.New()
	hash, err := argon2ID.GenerateFromPassword(newPassword)
	if err != nil {
		panic(err)
	}
	err = database.DBConn.Model(&models.User{}).Where("Username = ?", user.Username).Update("Password", hash).Error
	if err != nil {
		return err
	}
	return nil
}
