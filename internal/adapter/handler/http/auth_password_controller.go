package http

import (
	"better_auth/internal/service"

	"github.com/labstack/echo/v4"
)

type AuthPasswordController struct {
	service service.AuthPasswordService
}

func NewAuthPasswordController(authService service.AuthPasswordService) *AuthPasswordController {
	return &AuthPasswordController{
		service: authService,
	}
}

func (c *AuthPasswordController) Login(ctx echo.Context) error {
	// function ini diguankan untuk melakukan login

	// iisinya adalah validatsi input dan mengirimkan ke service

	// handling response error dan response sukksess
	return nil
}


func (c *AuthPasswordController) Register(ctx echo.Context) error {
	return nil
}

func (c *AuthPasswordController) Logout(ctx echo.Context) error {
	return nil
}

