package http

import "github.com/labstack/echo/v4"

type Router struct {
	authPasswordCtrl *AuthPasswordController
}

func NewRouter(
	authPasswordCtrl *AuthPasswordController,
) *Router {
	return &Router{
		authPasswordCtrl: authPasswordCtrl,
	}
}

func (rt *Router) RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/v1")

	authGroup := api.Group("/auth")
	{
		authGroup.POST("/login", rt.authPasswordCtrl.Login)
	}
}
