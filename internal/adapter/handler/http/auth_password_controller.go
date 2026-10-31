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
	return nil
}
