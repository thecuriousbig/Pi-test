package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	App      app
	Endpoint endpoint
	Redis    redis
	Database database
	JWT      jwt
}

type app struct {
	EnableSwagger bool `envconfig:"ENABLE_SWAGGER" default:"true"`
}

type endpoint struct {
	Port string `envconfig:"PORT" default:"8080"`
}

type redis struct {
	Host     string `envconfig:"REDIS_HOST" default:"localhost"`
	Port     string `envconfig:"REDIS_PORT" default:"6379"`
	Password string `envconfig:"REDIS_PASSWORD"`
}

type database struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	Username string `envconfig:"DB_USERNAME" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD"`
	Database string `envconfig:"DB_NAME" default:"pi"`
}

type jwt struct {
	Secret          string `envconfig:"JWT_SECRET"`
	AUD             string `envconfig:"JWT_AUD"`
	ISS             string `envconfig:"JWT_ISS"`
	ExpiresHours    uint   `envconfig:"JWT_EXPIRES_HOURS" default:"730"`
	AutoLogoffHours uint   `envconfig:"JWT_AUTO_LOGOFF_HOURS" default:"730"`
}

var cfg config

func New() {
	_ = godotenv.Load()
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("read env error : %s", err.Error())
	}
}

func Get() config {
	return cfg
}
