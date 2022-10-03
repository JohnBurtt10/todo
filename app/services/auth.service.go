package services

import (
	// "github.com/gofiber/fiber/v2/middleware/basicauth"

	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error         { return nil }
func Logout(ctx *fiber.Ctx) error        { return nil }
func Signup(ctx *fiber.Ctx) error        { return nil }
func ResetPassword(ctx *fiber.Ctx) error { return nil }

// // Or extend your config for customization
// app.Use(basicauth.New(basicauth.Config{
// 	Users: map[string]string{},
// 	Realm: "Forbidden",
// 	Authorizer: func(user, pass string) bool {
// 		var test User
// 		if err := db.Where("Username = ?", user).First(&test).Error; err != nil {
// 			// Username does not match any accounts in database
// 			if errors.Is(err, gorm.ErrRecordNotFound) {
// 				fmt.Println("no record found")
// 			}
// 			return false
// 		}
// 		// Password entered matches that of the password on record for account for username entered
// 		if test.Pass == pass {

// 			return true
// 		}
// 		// Password entered does not match that of the password on record for account for username entered
// 		fmt.Println("password doens't match")
// 		return false
// 	},
// 	// Unauthorized: func(c *fiber.Ctx) error {
// 	// 	return c.SendFile("./unauthorized.html")
// 	// },
// 	ContextUsername: "_user",
// 	ContextPassword: "_pass",
// }))
