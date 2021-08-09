package pkg

import (
	"context"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func Validate(ctx context.Context, request interface{}) error {
	value := reflect.ValueOf(request)
	validate := validator.New()
	switch value.Kind() {
	case reflect.Ptr:
		return validate.Struct(value.Elem().Interface())
	case reflect.Struct:
		return validate.Struct(value)
	}
	return nil
}
