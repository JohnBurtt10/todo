package routes

import (
	apiControllers "github.com/JohnBurtt10/go/app/controllers/api"
	"github.com/JohnBurtt10/go/app/middleware"
	"github.com/gofiber/fiber/v2"
)

func TaskRoute(route fiber.Router) {
	route.Get("", middleware.RequireSession, apiControllers.GetTasks)
	route.Post("", middleware.RequireSession, apiControllers.CreateTask)
	route.Patch("/:id", middleware.RequireSession, apiControllers.UpdateTask)
	route.Delete("/:id", middleware.RequireSession, apiControllers.DeleteTask)
}
