package ruang

import (
	"Diskusiaza-BE/app/http/middleware"
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/constants"
	"Diskusiaza-BE/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func createRuangController(c echo.Context) error {
	judul := c.FormValue("judul")
	deskripsi := c.FormValue("deskripsi")
	userID := c.Get("user_id").(int)
	ruang := model.Ruang{
		Judul:     judul,
		Deskripsi: deskripsi,
		UserID:    userID,
	}

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	user := middleware.GetDataFromToken(token)
	userID = int(user["id"].(float64))

	err := database.DB.Create(&ruang).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "proses create ruang gagal",
		})
	}
	database.DB.Create(&ruang)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "proses create ruang berhasil",
		"ruang":   ruang,
	})

}

func getRuangController(c echo.Context) error {
	ruang := model.Ruang{}
	database.DB.Find(&ruang)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "proses get ruang berhasil",
		"ruang":   ruang,
	})
}

func EditRuangController(c echo.Context) error {

	judul := c.FormValue("judul")
	deskripsi := c.FormValue("deskripsi")
	userID := c.Get("user_id").(int)
	ruang := model.Ruang{
		Judul:     judul,
		Deskripsi: deskripsi,
		UserID:    userID,
	}

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	user := middleware.GetDataFromToken(token)
	userID = int(user["id"].(float64))

	err := database.DB.Save(&ruang).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "proses edit ruang gagal",
		})
	}
	database.DB.Save(&ruang)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "proses edit ruang berhasil",
		"ruang":   ruang,
	})
}

func deleteRuangController(c echo.Context) error {
	ruang := model.Ruang{}
	_ = c.Bind(&ruang)

	database.DB.Delete(&ruang)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "proses delete ruang berhasil",
	})
}
