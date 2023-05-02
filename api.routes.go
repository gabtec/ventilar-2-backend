package main

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes is a function that exports all API route endpoints
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	// healthCheck
	api.Get("/healthcheck", ControllerHealthCheck)

	// // auth
	// api.Post("/login", controllers.Login)

	// // users
	// api.Post("/users", middleware.Protected(), controllers.CreateUser)
	// api.Get("/users", middleware.Protected(), controllers.GetUsers)
	// api.Get("/users/:id", middleware.Protected(), controllers.GetUserByID)
	// api.Patch("/users/:id", middleware.Protected(), controllers.UpdateUserPassword)
	
	// wards
	api.Get("/wards", GetWardsHandler)
	// api.Post("/wards", middleware.Protected(), controllers.CreateWard)
	// api.Get("/wards/:id", middleware.Protected(), controllers.GetWardByID)
}