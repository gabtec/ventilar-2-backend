package main

import (
	"log"

	"github.com/gabtec/ventilar-2-backend/types"
	"github.com/gofiber/fiber/v2"
)

// GetWardsHandler - GET /api/wards/ endpoint handler.
func GetWardsHandler(c *fiber.Ctx) error {
	wards := []types.Ward{}

	// err := DbInstance.Model(&Ward{}).Preload("Users").Find(&wards).Error
	err := DbInstance.Model(&types.Ward{}).Find(&wards).Error
	if err != nil {
		log.Fatal(err.Error())
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid resource data.", "error": err.Error()})
	}

	return c.Status(200).JSON(wards)
}

// PostWardHandler - POST /api/wards/ endpoint handler.
func PostWardHandler(c *fiber.Ctx) error {
	// get body
	ward := new(types.Ward)
	if err := c.BodyParser(ward); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// validation
	// err := validate.Struct(ward)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// 	return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid resource data.", "error": err.Error()})
	// }

	DbInstance.Create(&ward)
	return c.Status(201).JSON(ward)
}
