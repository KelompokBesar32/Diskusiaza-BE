package therad

import (
	"Diskusiaza-BE/app/http/middleware"
	"Diskusiaza-BE/app/repository/therad"
	"Diskusiaza-BE/constants"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Therad general = all function can use management therad_normal(pertanyaan) or therad_ruang(in ruang)

func DeleteTheradController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("therad_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameters therad_id",
		})
	}
	therad.DeleteTherad(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "therad successfully deleted",
	})
}

func GetAllTheradController(c echo.Context) error {
	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	res := therad.GetAllListTherad(int(middleware.GetDataFromToken(token)["id"].(float64)))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}

func GetByIdTheradController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("therad_id"))
	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameters therad_id",
		})
	}
	res := therad.GetTheradById(int(middleware.GetDataFromToken(token)["id"].(float64)), id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}

func GetSearchTheradController(c echo.Context) error {
	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]

	key := c.QueryParam("key")
	if key == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Key can't empty",
		})
	}

	res := therad.GetSearchTheradByJudulAndIsi(key, int(middleware.GetDataFromToken(token)["id"].(float64)))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
		"key":  key,
	})
}

func GetTheradByKategoriIdController(c echo.Context) error {
	kategoriId, err := strconv.Atoi(c.Param("kategori_id"))
	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameters kategori_id",
		})
	}
	res := therad.GetTheradByKategoriId(kategoriId, int(middleware.GetDataFromToken(token)["id"].(float64)))
	kategori := therad.GetKategoriTheradById(kategoriId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":          res,
		"kategori_name": kategori.KategoriName,
	})
}

func GetTrendingTheradController(c echo.Context) error {
	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	res := therad.GetTrendingTherad(int(middleware.GetDataFromToken(token)["id"].(float64)))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    res,
		"message": "favorite therad by total_like",
	})
}
