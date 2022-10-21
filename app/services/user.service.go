package services

import (
	"errors"
	"log"

	"github.com/JohnBurtt10/go/app/models"
	"github.com/JohnBurtt10/go/app/repos"
	"github.com/JohnBurtt10/go/database"
	"github.com/gofiber/fiber/v2"
)

func SessionUserID(c *fiber.Ctx) (uint, error) {
	sess, err := database.SessionStore.Get(c)
	if err != nil {
		return 0, errors.New("Session not found")
	}
	sessionUser := sess.Get("ID")
	if sessionUser == nil {
		return 0, errors.New("User not found in session")
	}
	t := sessionUser.(uint)

	return t, nil
}

func UserTemplateFromContext(c *fiber.Ctx) (*models.User, error) {
	userID, err := SessionUserID(c)
	if err != nil {
		return nil, err
	}
	u := new(models.User)
	err = repos.FindUserByID(u, userID)
	if err != nil {
		log.Printf("error could not find user (%s)\n", err.Error())
		return nil, errors.New("invalid user")
	}
	return u, nil
}
