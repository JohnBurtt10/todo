package main

import (
	"fmt"

	"github.com/JohnBurtt10/go/app/models"
	configuration "github.com/JohnBurtt10/go/config"
	"github.com/JohnBurtt10/go/database"
	"github.com/JohnBurtt10/go/routes"
	apiRoutes "github.com/JohnBurtt10/go/routes/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
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
	engine := html.New("./views", ".html")

	app := App{
		App: fiber.New(fiber.Config{
			Views: engine,
		}),
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

	// setupBasicAuth(app.App)
	setupRoutes(app.App)

	return app
}

//sess.Set("name", "john")

func setupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		sess, err := database.SessionStore.Get(c)
		if err != nil {
			panic(err)
		}
		if sess.Get("user") == "" {
			return c.Status(fiber.StatusTemporaryRedirect).Redirect("/signup")
		}
		// Get value

		// Set key/value
		sess.Set("user", uint(1))
		sess.Set("ID", uint(1))
		// keys := sess.Keys()

		// Save session
		if err := sess.Save(); err != nil {
			panic(err)
		}

		var user models.User
		database.DBConn.First(&user, 1)
		var tasks []models.Task
		database.DBConn.Model(&user).Association("Tasks").Find(&tasks)

		return c.Render("basictemplating", fiber.Map{
			"Title": tasks,
		})
	})

	app.Get("/signup", func(c *fiber.Ctx) error {
		if err != nil {
			return c.Render("auth/signup", fiber.Map{
				"Name": "Sign In",
			})
		} else {
			return c.Status(fiber.StatusTemporaryRedirect).Redirect("/")
		}
		// return c.Status(fiber.StatusOK).JSON(fiber.Map{
		// 	"success": true,
		// 	"message": "You are at the signup endpoint 😉",
		// })
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the login endpoint 😉",
		})
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
		sess, err := database.SessionStore.Get(c)
		if err != nil {
			panic(err)
		}
		// Get value

		// Set key/value
		sess.Set("user", uint(1))
		// keys := sess.Keys()

		// Save session
		if err := sess.Save(); err != nil {
			panic(err)
		}
		// fmt.Println(keys)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})

	apiRoutes.TaskRoute(api.Group("/tasks"))

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
