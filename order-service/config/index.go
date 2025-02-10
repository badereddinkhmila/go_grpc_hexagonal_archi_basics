package config

import (
	"errors"
	"io/fs"
	"log"
	"strings"
	"sync"

	viper "github.com/spf13/viper"
)

/*============================================================================*/
/*=====*                            Accessor                            *=====*/
/*============================================================================*/

func Environment() string { return viper.GetString("ENVIRONMENT") }

func IsProduction() bool { return Environment() == "production" }

func IsStaging() bool { return Environment() == "staging" }

func IsDemo() bool { return Environment() == "demo" }

func IsDevelopment() bool { return Environment() == "development" }

func IsOnline() bool { return IsProduction() || IsStaging() || IsDemo() }

func IsTest() bool { return Environment() == "test" }

func APP() app {
	setup().appOnce.Do(func() { setup().appConfig.load() })
	return setup().appConfig
}

func PostgreSQL() postgreSQL {
	setup().pgOnce.Do(func() { setup().pgConfig.load() })
	return setup().pgConfig
}

/*============================================================================*/
/*=====*                            Container                           *=====*/
/*============================================================================*/

var once sync.Once
var config *container

type container struct {
	// App
	appOnce   sync.Once
	appConfig app

	// PostgreSQL
	pgOnce   sync.Once
	pgConfig postgreSQL
}

func setup() *container {
	once.Do(func() {
		viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

		// Load from .env
		viper.SetConfigType("env")
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			// Do nothing
		} else if errors.Is(err, fs.ErrNotExist) {
			// Do nothing
		} else if err != nil {
			log.Fatalf("Fatal error loading .env: %v", err)
		}

		// Load from .env.local
		viper.SetConfigType("env")
		viper.SetConfigFile(".env.local")
		err = viper.MergeInConfig()
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			// Do nothing
		} else if errors.Is(err, fs.ErrNotExist) {
			// Do nothing
		} else if err != nil {
			log.Fatalf("Fatal error loading .env.local: %v", err)
		}

		// Populate config
		config = &container{}
	})
	return config
}
