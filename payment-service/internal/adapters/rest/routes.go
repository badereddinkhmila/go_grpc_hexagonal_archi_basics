package rest

import (
	"time"

	fiber "github.com/gofiber/fiber/v2"
	compress "github.com/gofiber/fiber/v2/middleware/compress"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
	requestid "github.com/gofiber/fiber/v2/middleware/requestid"
)

func Routes() *fiber.App {
	// New Fiber app
	router := fiber.New()
	router.Server().ReadTimeout = 5 * time.Minute
	router.Server().IdleTimeout = 1 * time.Minute
	router.Server().WriteTimeout = 5 * time.Minute

	// Set Middlewares
	router.Use(requestid.New())
	router.Use(compress.New(compress.Config{
		Level: compress.Level(4),
	}))
	router.Use(cors.New())

	router.Get("/rest/api/v1", func(c *fiber.Ctx) error {
		return c.JSON("Hello from server")
	})
	return router
}
