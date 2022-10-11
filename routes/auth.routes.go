package routes

import (
	"github.com/JohnBurtt10/go/app/services"

	"github.com/gofiber/fiber/v2"
)

// AuthRoutes containes all the auth routes
func AuthRoutes(route fiber.Router) {
	// route.Get("/resetPassword/:token", services.GetResetPassword)
	// route.Post("/resetPassword", services.CreatePasswordReset)
	route.Patch("/resetPassword", services.ResetPassword)
	route.Post("/logout", services.Logout)
	route.Post("/signup", services.Signup)
	route.Post("/login", services.Login)
}
