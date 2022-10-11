package services

import (
	"errors"

	"github.com/JohnBurtt10/go/database"
	"github.com/gofiber/fiber/v2"
)

func SessionUserID(c *fiber.Ctx) (uint, error) {
	sess, err := database.SessionStore.Get(c)
	if err != nil {
		return 0, errors.New("Session not found")
	}
	sessionUser := sess.Get("user")
	if sessionUser == nil {
		return 0, errors.New("User not found in session")
	}
	t := sessionUser.(uint)

	return t, nil
}
