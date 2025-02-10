package rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type RestAdapter struct {
	api  *fiber.App
	port int
}

func NewRestAdapter(api *fiber.App, port int) *RestAdapter {
	return &RestAdapter{api: api, port: port}
}

func (ra *RestAdapter) Run() {
	ra.api.Listen(fmt.Sprintf(":%d", ra.port))
	fmt.Println("[Rest Server] Fiber api started successfully")
}
