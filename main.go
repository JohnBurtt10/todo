package main

import (
	"fmt"
	"log"

	"github.com/JohnBurtt10/go/app/services"
	configuration "github.com/JohnBurtt10/go/config"
	"github.com/JohnBurtt10/go/database"
	"github.com/JohnBurtt10/go/routes"
	apiRoutes "github.com/JohnBurtt10/go/routes/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/*
docker exec -it my-postgres bash
psql mytestdb postgres
select * from friends.test
*/

type App struct {
	*fiber.App

	DB      *database.Database
	Session *session.Store
}

var db *gorm.DB
var err error

func setupApp(config *configuration.Config) App {
	database.Setup()

	app := App{
		App:     fiber.New(*config.GetFiberConfig()),
		Session: session.New(config.GetSessionConfig()),
	}

	app.DB = (&database.Database{
		DB: database.DBConn,
	})

	database.SessionStore = app.Session
	app.Session.RegisterType("")
	var typeUint uint = 1
	app.Session.RegisterType(typeUint)
	var typeBool bool = false
	app.Session.RegisterType(typeBool)

	setupRoutes(app.App)

	return app
}

// TODO: add success and message fiber app fields
func setupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		sess, err := database.SessionStore.Get(c)
		if err != nil {
			panic(err)
		}
		fmt.Println("Username:", sess.Get("username"))
		if sess.Get("username") == nil {
			return c.Status(fiber.StatusTemporaryRedirect).Redirect("/signup")
		}

		user, err := services.UserTemplateFromContext(c)
		if err != nil {
			log.Printf("error could not find user (%s)\n", err.Error())
			if err := sess.Destroy(); err != nil {
				log.Print("main/setupRoutes error destroying session", err.Error())
			}
			return fiber.NewError(fiber.StatusInternalServerError, "session user not found")
		}

		return c.Render("dashboard", fiber.Map{
			"User":        user,
			"SettingIcon": user.Firstname[0:1],
		})
	})
	app.Static("/", "./static/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    false,
		Index:     "/",
	})

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("signup", fiber.Map{})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})

	})

	app.Get("/changepassword", func(c *fiber.Ctx) error {
		return c.Render("changepassword", fiber.Map{})
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the logout endpoint 😉",
		})
	})

	// Auth Group
	authGroup := app.Group("/auth")

	routes.AuthRoutes(authGroup)

	// api group
	api := app.Group("/api")

	// give response when at /api
	api.Get("", func(c *fiber.Ctx) error {

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})

	apiRoutes.TaskRoute(api.Group("/tasks"))
	apiRoutes.UserRoute(api.Group("/user"))

	// 404 - last route
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendFile("./static/private/404.html")
	})
}

func main() {
	config := configuration.GetInstance()
	app := setupApp(config)

	app.Listen(":4000")
	fmt.Println("Successfully connected!")

}
