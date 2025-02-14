package rest

import (
	"context"
	"log"
	"product-service/internal/application/core/api"
	"product-service/internal/application/core/dto"

	"github.com/gofiber/fiber/v2"
)

type RestHandler struct {
	api *api.Application
}

func NewRestHandler(api api.Application) *RestHandler {
	return &RestHandler{api: &api}
}

func (rh *RestHandler) CreateProduct(c *fiber.Ctx) error {
	cpDTO := dto.CreateProductDTO{}
	err := c.BodyParser(cpDTO)
	if err != nil {
		log.Fatalf("Error parsing request: %v", err)
		return err
	}
	ctx := context.Background()
	product, err := rh.api.CreateProduct(ctx, cpDTO)
	if err != nil {
		log.Fatalf("Rest CreateProduct: %v", err)
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(product)
}
