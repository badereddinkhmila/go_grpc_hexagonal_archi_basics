package rest

import (
	"product-service/internal/ports"
)

type RestAdapter struct {
	rest ports.RestPort
	port int
}

func NewRestAdapter(rest ports.RestPort, port int) *RestAdapter {
	return &RestAdapter{rest: rest, port: port}
}
