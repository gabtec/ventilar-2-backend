package main

import (
	"github.com/gofiber/fiber/v2"
)

// ControllerHealthCheck - endpoint to check api health status is a hanlder function for GET /api/auth/login.
func ControllerHealthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"health": fiber.Map{
			"status":       "OK",
			"running":      true,
			"apiType":      "REST",
			"dependencies": "none",
			"database": fiber.Map{
				"connection":   "none",
				"responseTime": "0",
			},
			"memory":        "ok",
			"messageQueues": "no",
		},
	})
}
