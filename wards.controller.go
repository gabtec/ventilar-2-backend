package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// GetWardsHandler - GET /api/wards/ endpoint handler
func GetWardsHandler(c *fiber.Ctx) error {

	wards := []Ward{}

	err := DbInstance.Model(&Ward{}).Preload("Users").Find(&wards).Error

	if err != nil {
		log.Fatal(err.Error())
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid resource data.", "error": err.Error()})
	}

	return c.Status(200).JSON(wards)
}
