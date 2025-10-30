package unit

import (
	"testing"
	"workflowmanager/app/components/services"
	"workflowmanager/app/models"
	"workflowmanager/app/util"
)

func TestAuthService_RegisterUsingCredentials(test *testing.T) {
	const configFilePath string = "../../app/resources/dev_env.yaml"
	_ = util.LoadConfigs(configFilePath)
	userRepository := NewMockUserRepository()
	userService := services.NewAuthService(userRepository, nil)
	registerRequest := models.RegisterRequest{
		Email:    "test",
		Password: "test",
		Username: "test",
	}
	expectedUser := models.User{
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
		Username: registerRequest.Username,
	}
	actualUser, err := userService.RegisterUsingCredentials(registerRequest)
	if err != nil {
		test.Errorf("the issue during register user")
	}
	if err := util.VerifyPassword(actualUser.Password, registerRequest.Password); err != nil ||
		expectedUser.Username != actualUser.Username ||
		expectedUser.Email != actualUser.Email {
		test.Errorf("the expected user and actual user are different \n"+
			"expected user email: %s => actual user email: %s \n"+
			"expected username: %s => actual username: %s \n"+
			"password verify: %v",
			expectedUser.Email, actualUser.Email, expectedUser.Username, actualUser.Username, err)
	}
}
