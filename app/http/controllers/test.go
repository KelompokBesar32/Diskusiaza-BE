package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func TestController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Capstone Project API",
	})
}
