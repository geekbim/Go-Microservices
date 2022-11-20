package valueobject

import "errors"

type AppTypeEnum string

const (
	APP_TYPE_A AppTypeEnum = "A"
	APP_TYPE_B AppTypeEnum = "B"
)

type AppType struct {
	value AppTypeEnum
}

func NewAppType(value AppTypeEnum) AppType {
	return AppType{value: value}
}

func StringToAppType(str string) (AppTypeEnum, error) {
	var value AppTypeEnum
	switch str {
	case "A":
		value = APP_TYPE_A
	case "B":
		value = APP_TYPE_B
	default:
		return value, errors.New("app type not found")
	}
	return value, nil
}

func (a *AppType) GetValue() AppTypeEnum {
	return a.value
}
