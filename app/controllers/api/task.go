package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/JohnBurtt10/go/app/models"
	"github.com/JohnBurtt10/go/app/repos"
	"github.com/JohnBurtt10/go/app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func GetTasks(c *fiber.Ctx) error {
	UserID, err := services.SessionUserID(c)
	if err != nil {
		log.Fatal(err)
	}

	tasks, err := repos.GetTasks(UserID)
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

}

func UpdateTask(c *fiber.Ctx) error {
	var body models.Task

	if err := c.BodyParser(&body); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}
	aTask, err := repos.UpdateTask(body.ID, body.TaskName, body.Assignee, body.IsDone)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	result := &models.TaskResponse{}
	if err := copier.Copy(&result, &aTask); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot map results",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    result,
	})

}

func DeleteTask(c *fiber.Ctx) error {
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
	UserID, err := services.SessionUserID(c)
	if err != nil {
		log.Fatal(err)
	}
	var body models.Task

	if err := c.BodyParser(&body); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}
	aTask, err := repos.CreateTask(body.TaskName, body.Assignee, body.IsDone, UserID)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	result := &models.TaskResponse{}
	if err := copier.Copy(&result, &aTask); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot map results",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    result,
	})
}
