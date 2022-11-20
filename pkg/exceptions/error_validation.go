package exceptions

import (
	"errors"
	"fmt"
)

func ErrorRequired(field string) error {
	return errors.New(fmt.Sprintf("%s is required", field))
}

func ErrorInvalid(field string) error {
	return errors.New(fmt.Sprintf("%s is invalid", field))
}
