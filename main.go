package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/sirakytk/api-with-go/databases"
	"github.com/sirakytk/api-with-go/routes"
)

func main() {
	//konesi ke db
	databases.InitializeDB()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Set routes
	routes.Routes(app)

	app.Listen(":8000")
}
