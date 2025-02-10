package config

import (
	"fmt"
	"log"

	validate "shipping-service/config/validator"

	viper "github.com/spf13/viper"
)

type consul struct {
	ID   string `validate:"required"`
	Host string `validate:"required"`
	Port int    `validate:"required"`
}

func (consul) namespace() string         { return "CONSUL" }
func (obj consul) key(key string) string { return fmt.Sprintf("%s_%s", obj.namespace(), key) }

func (obj *consul) load() {
	obj.Host = viper.GetString(obj.key("Host"))
	obj.Port = viper.GetInt(obj.key("Port"))
	obj.ID = viper.GetString(obj.key("ID"))

	// Validate struct
	if err := validate.ValidateStruct(obj); err != nil {
		log.Fatal(err)
	}
}
