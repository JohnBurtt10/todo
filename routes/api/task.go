package routes

import (
	apiControllers "github.com/JohnBurtt10/go/app/controllers/api"
	"github.com/gofiber/fiber/v2"
)

// func TaskRoute(route fiber.Router) {
// 	route.Get("", middleware.RequirePermissions([]string{"tasks:read"}), apiControllers.GetTasks)
// 	// route.Get("/latestTimeStamp", middleware.RequirePermissions([]string{"tasks:read"}), apiControllers.GetTasksLatestTimeStamp)
// 	// route.Get("/assignees", middleware.RequirePermissions([]string{"tasks:read"}), apiControllers.SearchTaskAssignees)
// 	// route.Get("/:id", middleware.RequirePermissions([]string{"tasks:read"}), apiControllers.GetTask)
// 	route.Post("", middleware.RequirePermissions([]string{"tasks:create"}), apiControllers.CreateTask)
// 	// route.Patch("/:id", middleware.RequirePermissions([]string{"tasks:update"}), apiControllers.UpdateTask) IMPLEMENT THIS
// 	route.Delete("/:id", middleware.RequirePermissions([]string{"tasks:delete"}), apiControllers.DeleteTask)
// }

func TaskRoute(route fiber.Router) {
	route.Get("", apiControllers.GetTasks)
	// route.Get("/latestTimeStamp", middleware.RequirePermissions([]string{"tasks:read"}), apiControllers.GetTasksLatestTimeStamp)
	// route.Get("/assignees", middleware.RequirePermissions([]string{"tasks:read"}), apiControllers.SearchTaskAssignees)
	// route.Get("/:id", middleware.RequirePermissions([]string{"tasks:read"}), apiControllers.GetTask)
	route.Post("", apiControllers.CreateTask)
	// route.Patch("/:id", middleware.RequirePermissions([]string{"tasks:update"}), apiControllers.UpdateTask) IMPLEMENT THIS
	route.Delete("/:id", apiControllers.DeleteTask)
}
