package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "OriginLab API (v3)",
	})

	// Global middlewares
	app.Use(logger.New())
	app.Use(recover.New())

	// Built-in Fiber v3 cloud health check probes (/livez, /readyz, /startupz)
	app.Get(healthcheck.LivenessEndpoint, healthcheck.New())
	app.Get(healthcheck.ReadinessEndpoint, healthcheck.New())
	app.Get(healthcheck.StartupEndpoint, healthcheck.New())

	// Application-level health endpoint
	app.Get("/healthz", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "ok",
			"service": "originlab-api",
			"fiber":   "v3",
		})
	})

	// API v1 route group placeholder
	api := app.Group("/api/v1")
	api.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "healthy",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting OriginLab API on :%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
