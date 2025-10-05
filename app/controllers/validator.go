package controllers

import (
	"errors"
	"workflowmanager/app/models"
)

func validateAuthRequestBody(registerRequest models.RegisterRequest) error {
	if registerRequest.Email == "" ||
		registerRequest.Username == "" ||
		registerRequest.Password == "" {
		return errors.New("email, username, and password are required")
	}
	return nil
}
