package ruang

import (
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/database"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func createRuang(c echo.Context) error {
	validate := validator.New()
	ruang := model.Ruang{}
	_ = c.Bind(&ruang)

	if err := validate.Struct(ruang); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "ada inputan yang tidak sesuai",
		})
	}

	database.DB.Create(&ruang)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "proses create ruang berhasil",
	})
}

func getRuang(c echo.Context) error {
	ruang := model.Ruang{}
	database.DB.Find(&ruang)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"ruang": ruang,
	})
}

func putRuang(c echo.Context) error {
	validate := validator.New()
	ruang := model.Ruang{}
	_ = c.Bind(&ruang)

	if err := validate.Struct(ruang); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "ada inputan yang tidak sesuai",
		})
	}

	database.DB.Save(&ruang)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "proses update ruang berhasil",
	})
}
