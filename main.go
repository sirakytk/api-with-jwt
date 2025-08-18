package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirakytk/api-with-go/databases"
	"github.com/sirakytk/api-with-go/routes"
)

func main() {
	//konesi ke db
	databases.InitializeDB()

	app := fiber.New()

	// Set routes
	routes.Routes(app)

	app.Listen(":8000")
}
