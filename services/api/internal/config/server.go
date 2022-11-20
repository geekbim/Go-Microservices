package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/hashicorp/go-multierror"
)

type ApiConfig struct {
	AppPort     string
	GrpcPort    string
	ServicePort Port
}

type Port struct {
	Auth string
	User string
}

func Api() (*ApiConfig, *multierror.Error) {
	var multierr *multierror.Error

	c := &ApiConfig{
		AppPort:  os.Getenv("APP_PORT"),
		GrpcPort: os.Getenv("GRPC_PORT"),
		ServicePort: Port{
			Auth: os.Getenv("AUTH_PORT"),
			User: os.Getenv("USER_PORT"),
		},
	}

	if err := c.required("APP_PORT", c.AppPort); err != nil {
		multierr = multierror.Append(multierr, err)
	}
	if err := c.required("GRPC_PORT", c.GrpcPort); err != nil {
		multierr = multierror.Append(multierr, err)
	}
	if err := c.required("AUTH_PORT", c.ServicePort.Auth); err != nil {
		multierr = multierror.Append(multierr, err)
	}
	if err := c.required("USER_PORT", c.ServicePort.User); err != nil {
		multierr = multierror.Append(multierr, err)
	}

	return c, multierr
}

func (c *ApiConfig) required(key string, value string) error {
	if value == "" {
		errorMsg := fmt.Sprintf("config %s is required", key)
		return errors.New(errorMsg)
	}
	return nil
}
