package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// RouteNotFoundErrorHandler - will catch 404 errors
func RouteNotFoundErrorHandler(ctx *fiber.Ctx) error {
	// If request reaches here, it's because no route matched

	// return fiber.NewError(404, "Route Not Found!")
	return ctx.Status(404).JSON(fiber.Map{
		"status": 404,
		"message": "Route Not Found!",
	})
}

// CustomDefaultErrorHandler - will catch all other errors, and customize them
func CustomDefaultErrorHandler(c *fiber.Ctx, err error) error {

	// TODO: filter my custom errors
	
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
			code = e.Code
	}

	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	// Return status code with error message
	return c.Status(code).SendString(err.Error())
}