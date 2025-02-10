package config

import (
	"fmt"
	"log"

	validate "order-service/config/validator"

	viper "github.com/spf13/viper"
)

type postgreSQL struct {
	Host     string `validate:"required"`
	Port     int    `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Database string `validate:"required"`
}

func (postgreSQL) namespace() string         { return "POSTGRES" }
func (obj postgreSQL) key(key string) string { return fmt.Sprintf("%s_%s", obj.namespace(), key) }

func (obj *postgreSQL) load() {
	obj.Host = viper.GetString(obj.key("Host"))
	obj.Port = viper.GetInt(obj.key("Port"))
	obj.Database = viper.GetString(obj.key("DB"))
	obj.User = viper.GetString(obj.key("USER"))
	obj.Password = viper.GetString(obj.key("PASSWORD"))
	// Validate struct
	if err := validate.ValidateStruct(obj); err != nil {
		log.Fatal(err)
	}
}
