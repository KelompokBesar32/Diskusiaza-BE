package routes

import (
	"Diskusiaza-BE/app/http/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// application routes
	e.GET("/", controllers.TestController)
	// authentication
	e.POST("/auth/register", controllers.RegisterController)
	e.POST("/auth/login", controllers.LoginController)
	e.GET("/auth/logout", controllers.LogoutController)
	// create therad
	return e
}
