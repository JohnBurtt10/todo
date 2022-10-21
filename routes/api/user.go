package routes

import (
	apiControllers "github.com/JohnBurtt10/go/app/controllers/api"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(route fiber.Router) {
	route.Post("", apiControllers.CheckIfUserExists)

}
