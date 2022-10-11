package services

import (
	"errors"

	"gorm.io/gorm"

	"github.com/JohnBurtt10/go/app/models"
	"github.com/JohnBurtt10/go/app/repos"
	"github.com/JohnBurtt10/go/database"
	"github.com/gofiber/fiber/v2"
)

//TODO: make it so that these functions use repo functions

func Login(ctx *fiber.Ctx) error {
	b := new(models.User)
	if err := ctx.BodyParser(&b); err != nil {
		return err
	}
	u := new(models.User)
	err := database.DBConn.Where("Username = ? AND Password = ?", b.Username, b.Password).Take(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "Invalid username or password")
	}

	sess, err := database.SessionStore.Get(ctx)
	if err != nil {
		panic(err)
	}

	// Set key/value
	sess.Set("ID", u.ID)
	sess.Set("username", u.Username)

	// save session
	if err := sess.Save(); err != nil {
		panic(err)
	}

	return ctx.JSON(&models.UserResponse{
		ID:       u.ID,
		Username: u.Username,
	})
}

func Signup(ctx *fiber.Ctx) error {
	b := new(models.User)

	if err := ctx.BodyParser(&b); err != nil {
		return err
	}
	u := new(models.User)
	err := database.DBConn.Where("Username = ?", b.Username).Take(&u).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "Username is taken")
	}

	if err := database.DBConn.Create(&b).Error; err != nil {
		return err
	}

	// why do we return this as an error

	return ctx.JSON(&models.UserResponse{
		ID:       u.ID,
		Username: u.Username,
	})
}

func Logout(ctx *fiber.Ctx) error {
	sess, err := database.SessionStore.Get(ctx)
	if err != nil {
		panic(err)
	}

	// Destry session
	if err := sess.Destroy(); err != nil {
		panic(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "You are logged out 😉",
	})
}

func ResetPassword(ctx *fiber.Ctx) error {
	sess, err := database.SessionStore.Get(ctx)
	if err != nil {
		panic(err)
	}
	type Request struct {
		OldPassword string `json:"oldpassword" validate:"omitempty,min=8,max=20,alphanum"`
		NewPassword string `json:"newpassword" validate:"omitempty,min=8,max=20,alphanum"`
	}

	b := new(Request)
	if err := ctx.BodyParser(&b); err != nil {
		return err
	}

	u, err := repos.GetUserByID(sess.Get("ID").(uint))
	if err != nil {
		return err
	}

	if b.OldPassword != u.Password {
		return fiber.NewError(fiber.StatusConflict, "Password doesn't match records")
	}

	if err := repos.ChangeUserPassword(u, b.NewPassword); err != nil {
		return err
	}

	return nil
}
