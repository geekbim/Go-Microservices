package config

import (
	"os"
	"strconv"
	"time"
)

type UserConfig struct {
	Port       string
	Timeout    time.Duration
	GrpcClient GrpcClient
}

type GrpcClient struct {
	Auth string
}

func User() UserConfig {
	duration, _ := strconv.Atoi(os.Getenv("APP_TIMEOUT"))
	return UserConfig{
		Port:    os.Getenv("USER_PORT"),
		Timeout: time.Duration(duration) * time.Second,
		GrpcClient: GrpcClient{
			Auth: os.Getenv("AUTH_PORT"),
		},
	}
}
