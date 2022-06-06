package routes

import (
	"Diskusiaza-BE/app/http/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// application routes
	e.GET("/", controllers.TestController)

	return e
}
