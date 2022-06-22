package therad

import (
	"Diskusiaza-BE/app/repository/therad"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetKategoriTheradController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": therad.GetKategoriTherad(),
	})
}
