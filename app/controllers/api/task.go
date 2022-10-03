package controllers

import (
	"fmt"
	"strconv"

	"github.com/JohnBurtt10/go/app/models"
	"github.com/JohnBurtt10/go/app/repos"
	"github.com/JohnBurtt10/go/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func GetTasks(c *fiber.Ctx) error {
	// sess, err := database.SessionStore.Get(c)
	fmt.Println(c.AllParams())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	var id uint
	id = uint(1)
	// test := "0"
	// // convert parameter value string to int
	// if v, err := strconv.ParseUint(test, 10, 32); err == nil {
	// 	id = uint(v)
	// } else {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": "Cannot parse ID",
	// 	})
	// }

	tasks, err := repos.GetTasks(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Tasks not found",
		})
	}
	result := &[]models.TaskResponse{}
	if err := copier.Copy(&result, &tasks); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot map results",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    result,
	})
	// if Task not available

}

func DeleteTask(c *fiber.Ctx) error {
	fmt.Println(c.Params("id"))
	paramId := c.Params("id")
	var id uint
	// convert parameter value string to int
	if v, err := strconv.ParseUint(paramId, 10, 32); err == nil {
		id = uint(v)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse ID",
		})
	}

	if err := repos.DeleteTask(id); err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
		})
	}
	// if Task not available
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Task not found",
	})
}

func CreateTask(c *fiber.Ctx) error {
	task := models.Task{}
	if err := c.BodyParser(task); err != nil {
		return err
	}
	database.DBConn.Select("Title").Create(&task)
	return nil
	// ...
}
