package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Welcome(c *fiber.Ctx) error {
	// Render index.html dengan data
	return c.Render("home/index", fiber.Map{
		"Title": "Hello, World!",
	})
}

func LoginPage(c *fiber.Ctx) error {
	return c.Render("auth/login", fiber.Map{
		"title": "Login",
	})
}

func RegisterPage(c *fiber.Ctx) error {
	return c.Render("auth/register", fiber.Map{
		"title": "Register",
	})
}

func HomePage(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"title": "Home Page",
		"name":  "User",
	})
}
