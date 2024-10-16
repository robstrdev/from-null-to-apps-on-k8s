package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Serve static files from the "public" directory.
	// This is where the compiled frontend code will be.
	app.Static("/", "./public")

	// Example API endpoint, returns "Hello, World!".
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Health check endpoint, returns "OK" if the server is running.
	// Used by Kubernetes to check if everything is fine.
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Get the port from the environment variable "SERVER_PORT".
	port := os.Getenv("SERVER_PORT")
	log.Fatal(app.Listen(port))
}
