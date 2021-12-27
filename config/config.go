package config

import (
	"log"
	"time"

	"github.com/jinzhu/configor"
)

// Config represents application configuration which can also be set
// via environment variable during runtime.
type Config struct {
	App struct {
		APP struct {
			URL       string `default:"http://localhost" env:"APP_URL"`
			Port      uint   `default:"8080" env:"APP_PORT"`
			Timezone  string `default:"UTC" env:"APP_TIMEZONE"`
			DebugMode bool   `default:"false" env:"APP_DEBUG_MODE"`
		}
		ADMIN struct {
			URL       string `default:"http://localhost" env:"APP_URL"`
			Port      uint   `default:"8081" env:"ADMIN_PORT"`
			Timezone  string `default:"UTC" env:"APP_TIMEZONE"`
			DebugMode bool   `default:"false" env:"APP_DEBUG_MODE"`

			// Values of access-control-allow-origin in response header.
			ResponseHeaderAllowOrigins []string `env:"ADMIN_RESPONSE_HEADER_ALLOW_ORIGINS"`
		}
		DevMode bool `default:"false" env:"DEV"`
	}

	Database struct {
		Postgres struct {
			Host     string `default:"localhost" env:"DB_HOST"`
			User     string `default:"root" env:"DB_USER"`
			Password string `default:"" env:"DB_PASSWORD"`
			Port     string `default:"5432" env:"DB_PORT"`
			Name     string `env:"DB_NAME"`
		}
	}

	Log struct {
		Level  string `env:"LOG_LEVEL" default:"info"`
		Format string `env:"LOG_FORMAT" default:"console"`
	}
}

// LoadConfig loads the configuration from `.env` file in the same
// directory as the application and populate the Config accordingly.
func LoadConfig() (*Config, error) {
	var config Config

	err := loadEnv()
	if err == nil {
		log.Println("Env file loaded")
	}

	err = configor.
		New(&configor.Config{AutoReload: true, AutoReloadInterval: time.Minute}).
		Load(&config)

	if err != nil {
		log.Println(err)
		log.Fatal("Error loading config")
	}

	return &config, err
}

// LoadTestConfig ...
func LoadTestConfig() (*Config, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	return config, err
}
