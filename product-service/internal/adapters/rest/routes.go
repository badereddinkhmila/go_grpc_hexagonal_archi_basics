package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	compress "github.com/gofiber/fiber/v2/middleware/compress"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
	requestid "github.com/gofiber/fiber/v2/middleware/requestid"
)

func (ra *RestAdapter) InitServer() error {
	// Fiber server config
	config := fiber.Config{
		ReadTimeout:  5 * time.Minute,
		IdleTimeout:  1 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	}
	// New Fiber app
	app := fiber.New(config)
	// Set Middlewares
	app.Use(requestid.New()).Use(compress.New(compress.Config{
		Level: compress.Level(4),
	})).Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// http server health check
	app.Get("/check", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	// Rest Client Handlers
	rh := NewRestHandler(ra.api)

	// Initialize main route
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/products", func(c *fiber.Ctx) error {
		ctx := context.Background()
		products, err := ra.api.GetProducts(ctx)
		if err != nil {
			log.Fatalf("[Get Products Route]: %v", err)
			return err
		}
		return c.JSON(products)
	})

	v1.Post("/products", rh.CreateProduct)

	v1.Get("/categories", func(c *fiber.Ctx) error {
		ctx := context.Background()
		categories, err := ra.api.GetCategories(ctx)
		if err != nil {
			log.Fatalf("[Get Categories Route]: %v", err)
			return err
		}
		return c.JSON(categories)
	})

	app.Get("*", func(c *fiber.Ctx) error {
		return c.JSON(http.StatusNotFound)
	})

	return app.Listen(fmt.Sprintf(":%d", ra.port))
}
