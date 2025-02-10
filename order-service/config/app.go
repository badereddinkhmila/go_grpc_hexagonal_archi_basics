package config

import (
	"fmt"
	"log"

	validate "order-service/config/validator"

	viper "github.com/spf13/viper"
)

type app struct {
	Host string `validate:"required"`
	Port int    `validate:"required"`
}

func (app) namespace() string         { return "APP" }
func (obj app) key(key string) string { return fmt.Sprintf("%s_%s", obj.namespace(), key) }

func (obj *app) load() {
	obj.Host = viper.GetString(obj.key("Host"))
	obj.Port = viper.GetInt(obj.key("Port"))
	// Validate struct
	if err := validate.ValidateStruct(obj); err != nil {
		log.Fatal(err)
	}
}
