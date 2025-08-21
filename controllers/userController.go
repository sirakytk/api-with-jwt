package controllers

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/sirakytk/api-with-go/databases"
// 	"github.com/sirakytk/api-with-go/models"
// )

// func GetUser(c *fiber.Ctx) error {
// 	var user models.User
// 	userID := c.Locals("userID")

// 	if err := databases.DB.First(&user, userID).Error; err != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 			"message": "User not found",
// 		})
// 	}

// 	return c.JSON(user)
// }
