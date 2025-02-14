package rest

import (
	"product-service/internal/application/core/api"
)

type RestAdapter struct {
	api  api.Application
	port int
}

func NewRestAdapter(api api.Application, port int) *RestAdapter {
	return &RestAdapter{api: api, port: port}
}
