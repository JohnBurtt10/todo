package services

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/JohnBurtt10/go/app/models"
	"github.com/JohnBurtt10/go/app/repos"
	"github.com/JohnBurtt10/go/app/utils/password/argon2id"
	"github.com/JohnBurtt10/go/database"
	"github.com/gofiber/fiber/v2"
)

//TODO: make it so that these functions use repo functions

func Login(ctx *fiber.Ctx) error {
	argon2ID := argon2id.New()
	b := new(models.User)
	if err := ctx.BodyParser(&b); err != nil {
		return err
	}

	// not needed
	u := new(models.User)
	err := repos.FindUserByUsername(u, b.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "Invalid username")
	}
	err = argon2ID.ComparePasswordAndHash(u.Password, b.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "Invalid password")
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

	s := fmt.Sprintf("Username: %s, ID: %d has logged in.", u.Username, u.ID)
	log.Printf(s)

	return ctx.JSON(&models.UserResponse{
		ID:       u.ID,
		Username: u.Username,
	})
}

func Signup(ctx *fiber.Ctx) error {
	argon2ID := argon2id.New()
	b := new(models.User)

	if err := ctx.BodyParser(&b); err != nil {
		return err
	}
	// not needed
	u := new(models.User)
	err := repos.FindUserByUsername(u, b.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "Username is taken")
	}

	hash, err := argon2ID.GenerateFromPassword(b.Password)
	if err != nil {
		panic(err)
	}

	user := &models.User{
		Firstname: b.Firstname,
		Lastname:  b.Lastname,
		Username:  b.Username,
		Password:  hash,
	}

	if err := database.DBConn.Create(&user).Error; err != nil {
		return err
	}

	sess, err := database.SessionStore.Get(ctx)
	if err != nil {
		panic(err)
	}

	// Set key/value
	sess.Set("ID", user.ID)
	sess.Set("username", user.Username)

	// save session
	if err := sess.Save(); err != nil {
		panic(err)
	}

	s := fmt.Sprintf("Username: %s, ID: %d has signed up.", user.Username, user.ID)
	log.Printf(s)

	// why do we return this as an error

	return ctx.JSON(&models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
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
	argon2ID := argon2id.New()
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

	u := new(models.User)
	err = repos.FindUserByID(u, sess.Get("ID").(uint))
	if err != nil {
		return err
	}

	err = argon2ID.ComparePasswordAndHash(b.OldPassword, u.Password)
	if err != nil {
		return err
	}
	if err := repos.ChangeUserPassword(u, b.NewPassword); err != nil {
		return err
	}

	return nil
}
