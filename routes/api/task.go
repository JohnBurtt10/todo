package routes

import (
	apiControllers "github.com/JohnBurtt10/go/app/controllers/api"
	"github.com/gofiber/fiber/v2"
)

func TaskRoute(route fiber.Router) {
	route.Get("", apiControllers.GetTasks)
	route.Post("", apiControllers.CreateTask)
	route.Patch("/:id", apiControllers.UpdateTask)
	route.Delete("/:id", apiControllers.DeleteTask)
}
