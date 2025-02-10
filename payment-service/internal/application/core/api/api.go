package api

import "payment-service/internal/ports"

type Application struct {
	db ports.DbPort
}

func NewApplication(db ports.DbPort) *Application {
	return &Application{
		db: db,
	}
}
