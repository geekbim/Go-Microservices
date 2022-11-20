package config

import (
	"go-grpc-example/pkg/config"
	"os"
	"strconv"
	"time"
)

type AuthConfig struct {
	Port     string
	TimeOut  time.Duration
	UserPort string
}

func convertInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))

	return v
}

func Auth() AuthConfig {
	authConfig := AuthConfig{
		Port:     os.Getenv("AUTH_PORT"),
		TimeOut:  time.Duration(convertInt("APP_TIMEOUT")) * time.Second,
		UserPort: os.Getenv("USER_PORT"),
	}

	err := authConfig.Validate()
	if err != nil {
		panic(err)
	}

	return authConfig
}

func (c *AuthConfig) Validate() error {
	fields := []string{
		"USER_PORT",
		"AUTH_PORT",
	}

	for _, field := range fields {
		err := config.Required(field)
		if err != nil {
			return err
		}
	}

	return nil
}
