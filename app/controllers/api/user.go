package controllers

import (
	"fmt"

	"github.com/JohnBurtt10/go/app/models"
	"github.com/JohnBurtt10/go/app/repos"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func CheckIfUserExists(c *fiber.Ctx) error {
	type Request struct {
		Username string `json:"username" validate:"omitempty,min=5,max=16,alphanum"`
	}
	b := Request{}
	if err := c.BodyParser(&b); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}
	user := &[]models.User{}
	err := repos.FindUserByUsername(user, b.Username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": true,
			"message": "error finding user",
		})
	}
	result := &[]models.UserResponse{}
	if err := copier.Copy(&result, &user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot map results",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    result,
	})

	// // if Task not available

}
