package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// SetupFiber - will initiate a GoFiber API.
func SetupFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "VentilAR REST API v0.0.1",
		// EnablePrintRoutes: true, // lists all defined routes (sos)
		// Catch any unhandled error
		ErrorHandler: CustomDefaultErrorHandler,
	})

	app.Use(cors.New())

	SetupRoutes(app)

	// 404 | Not Found!
	app.Use(RouteNotFoundErrorHandler)

	return app
}
