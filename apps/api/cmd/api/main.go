package main

import (
	"context"
	"log"
	"time"

	"github.com/NatdaniKhe/OriginLab/apps/api/internal/platform/config"
	"github.com/NatdaniKhe/OriginLab/apps/api/internal/platform/database"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "OriginLab API (v3)",
	})

	config := config.LoadConfig()
	ctx := context.Background()

	// Global middlewares
	app.Use(logger.New())
	app.Use(recover.New())

	pool, err := database.NewPool(ctx, config.DatabaseConfig.DATABASE_URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	app.Get(healthcheck.LivenessEndpoint, healthcheck.New())

	// API v1 route group
	api := app.Group("/api/v1")
	api.Get("/health", func(c fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 2*time.Second)
		defer cancel()

		if err := pool.Ping(ctx); err != nil {
			log.Printf("health check: db ping failed: %v", err)
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"status":   "degraded",
				"database": "down",
			})
		}

		return c.JSON(fiber.Map{
			"status":   "healthy",
			"database": "ok",
		})
	})

	port := config.PORT

	log.Printf("Starting OriginLab API on :%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
