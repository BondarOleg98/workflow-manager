package unit

import (
	"github.com/oklog/ulid/v2"
	"testing"
	"workflowmanager/app/components/services"
	"workflowmanager/app/models"
	"workflowmanager/app/util"
)

var registerRequest = models.RegisterRequest{
	Email:    "test",
	Password: "test",
	Username: "test",
}

var loginRequest = models.LoginRequest{
	Email:    "test",
	Password: "test",
}

func prepareTestConfigs() {
	const configFilePath string = "../../app/resources/dev_env.yaml"
	_ = util.LoadConfigs(configFilePath)
}

func TestAuthServiceRegisterUsingCredentials(test *testing.T) {
	prepareTestConfigs()
	userRepository := NewMockUserRepository()
	authService := services.NewAuthService(userRepository, nil)
	expectedUser := models.User{
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
		Username: registerRequest.Username,
	}
	actualUser, err := authService.RegisterUsingCredentials(registerRequest)
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

func TestAuthServiceLoginUsingCredentials(test *testing.T) {
	prepareTestConfigs()
	userRepository := NewMockUserRepository()
	refreshTokenRepository := NewRefreshTokenRepository()
	authService := services.NewAuthService(userRepository, refreshTokenRepository)
	_, err := authService.RegisterUsingCredentials(registerRequest)
	if err != nil {
		test.Errorf("the issue during register user")
	}
	accessToken, refreshToken, err := authService.Login(loginRequest)
	if err != nil {
		test.Errorf("the issue during login user")
	}
	claims, err := authService.ValidateToken(accessToken)
	userIdStr, isContextKeyExist := claims["sub"].(string)
	if err != nil {
	  test.Errorf("the issue during validating token")
	}
	if !isContextKeyExist {
		test.Errorf("invalid token claims")
	}
	_, err = ulid.Parse(userIdStr)
	if err != nil {
		test.Errorf("invalid userId")
	}
	if refreshToken == "" {
		test.Errorf("the refresh token is empty")
	}
}

func TestAuthServiceRefreshAccessToken(test *testing.T) {
	prepareTestConfigs()
	userRepository := NewMockUserRepository()
	refreshTokenRepository := NewRefreshTokenRepository()
	authService := services.NewAuthService(userRepository, refreshTokenRepository)
	_, err := authService.RegisterUsingCredentials(registerRequest)
	if err != nil {
		test.Errorf("the issue during register user")
	}
	_, refreshToken, err := authService.Login(loginRequest)
	if err != nil {
		test.Errorf("the issue during login user")
	}
	token, err := refreshTokenRepository.GetRefreshToken(refreshToken)
	if err != nil {
		test.Errorf("the issue during getting the refresh token")
	}
	if token.Revoked {
		test.Errorf("the refresh token was not revoked")
	}
	_, err = authService.RefreshAccessToken(refreshToken)
	token, err = refreshTokenRepository.GetRefreshToken(refreshToken)
	if err != nil {
		test.Errorf("the issue during getting the refresh token")
	}
	if !token.Revoked {
		test.Errorf("the refresh token was not revoked")
	}
}
