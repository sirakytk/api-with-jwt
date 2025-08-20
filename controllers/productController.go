package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirakytk/api-with-go/databases"
	"github.com/sirakytk/api-with-go/models"
)

func GetAllProduct(c *fiber.Ctx) error {
	var product []models.Produk

	if err := databases.DB.Find(&product).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Produk not found",
		})
	}

	// Assuming you want to return all products
	databases.DB.Find(&product)
	if len(product) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Produk not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Get All Produk",
		"data":    product,
	})
}

func GetProductByID(c *fiber.Ctx) error {
	var product models.Produk
	productID := c.Params("id")

	if err := databases.DB.First(&product, productID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Produk not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Get Produk By ID",
		"data":    product,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Produk

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	if err := databases.DB.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create product",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Produk created successfully",
		"data":    product,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	var product models.Produk
	productID := c.Params("id")

	if err := databases.DB.First(&product, productID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Produk not found",
		})
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	if err := databases.DB.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update product",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Produk updated successfully",
		"data":    product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	var product models.Produk
	productID := c.Params("id")

	if err := databases.DB.First(&product, productID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Produk not found",
		})
	}

	if err := databases.DB.Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete product",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Produk deleted successfully",
	})
}
