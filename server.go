package main

import (
	"fmt"

	// "log"

	"github.com/JohnBurtt10/go/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/*
docker exec -it my-postgres bash
psql mytestdb postgres
select * from friends.test
*/

// type App struct {
// 	*fiber.App

// 	DB	*database.DataBase
// 	Session *session.Store
// }

type App struct {
	*fiber.App

	DB      *database.Database
	Session *session.Store
}

var db *gorm.DB
var err error

func setupApp() App {
	database.Setup()

	app := App{
		App:     fiber.New(),
		Session: session.New(),
	}

	app.DB = (&database.Database{
		DB: database.DBConn,
	})

	database.SessionStore = app.Session

	setupRoutes(app.App)

	return app
}

func setupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the app endpoint 😉",
		})
	})

	// Auth Group
	// authGroup := app.Group("/auth")

	// routes.AuthRoutes(authGroup)

	// api group
	// api := app.Group("/api")

	// // give response when at /api
	// api.Get("", func(c *fiber.Ctx) error {
	// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 		"success": true,
	// 		"message": "You are at the api endpoint 😉",
	// 	})
	// })

	// apiRoutes.TaskRoute(api.Group("/tasks"))

	// // 404 - last route
	// app.Use(func(c *fiber.Ctx) error {
	// 	return c.Status(fiber.StatusNotFound).SendFile("./static/private/404.html")
	// })
}

func main() {
	app := setupApp()
	// var item models.User
	// REDIS
	// BASIC AUTH
	// Or extend your config for customization
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			// if err := database.DBConn.Where("Username = ? AND Password = ?", user, pass).Find(&item).RowsAffected; err == 1 {
			// 	fmt.Println("Sucessfully logging in")
			// 	return true
			// }
			return false
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the app endpoint 😉",
		})
	})

	// // This is the fiber version
	app.Listen(":3000")
	fmt.Println("Successfully connected!")

}
