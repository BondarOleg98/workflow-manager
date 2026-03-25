package components

import (
	"github.com/go-playground/validator/v10"
	"log/slog"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

func GetValidatorInstance() *validator.Validate {
	slog.Info("Getting the request validator")
	return validate
}
