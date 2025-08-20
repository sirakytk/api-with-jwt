package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirakytk/api-with-go/controllers"
)

func Routes(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/user", controllers.GetUser)
	app.Get("/", controllers.Welcome)
	app.Get("/login", controllers.LoginPage)
	app.Get("/register", controllers.RegisterPage)
	app.Get("/home", controllers.HomePage)
}
