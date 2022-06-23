package therad

import (
	"Diskusiaza-BE/app/http/middleware"
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/app/repository/therad"
	"Diskusiaza-BE/constants"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// Therad normal = create pertanyaan -> diluar ruang yang ditentukan (random question)

func GetListTheradNormalByUsersController(c echo.Context) error {
	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	res := therad.GetListTheradByUserId(int(middleware.GetDataFromToken(token)["id"].(float64)))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}

func CreateTheradNormalController(c echo.Context) error {
	judul := c.FormValue("judul")
	isi := c.FormValue("isi")
	file, _ := c.FormFile("file")
	kategoriTheradId, err := strconv.Atoi(c.FormValue("kategori_therad_id"))

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid form",
		})
	}

	// if file != null
	if file != nil {
		pathFile := os.Getenv("BASE_URL") + constants.StaticFileTherad + file.Filename
		src, _ := file.Open()
		dir, _ := os.Getwd()
		locationFile := filepath.Join(dir, constants.DirFileTherad, file.Filename)
		dst, _ := os.OpenFile(locationFile, os.O_WRONLY|os.O_CREATE, 06666)
		defer func(dst *os.File) {
			_ = dst.Close()
		}(dst)
		_, _ = io.Copy(dst, src)

		mTherad := model.Therad{
			Judul:            judul,
			Isi:              isi,
			File:             pathFile,
			Dilihat:          0,
			Status:           "active",
			KategoriTheradID: kategoriTheradId,
			UserID:           int(middleware.GetDataFromToken(token)["id"].(float64)),
		}
		therad.CreateNormalTherad(mTherad)
	} else {
		mTherad := model.Therad{
			Judul:            judul,
			Isi:              isi,
			Dilihat:          0,
			Status:           "active",
			KategoriTheradID: kategoriTheradId,
			UserID:           int(middleware.GetDataFromToken(token)["id"].(float64)),
		}
		therad.CreateNormalTherad(mTherad)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Therad pertanyaan anda berhasil diposting",
	})
}

func UpdateTheradNormalController(c echo.Context) error {
	judul := c.FormValue("judul")
	isi := c.FormValue("isi")
	file, _ := c.FormFile("file")
	kategoriTheradId, err := strconv.Atoi(c.FormValue("kategori_therad_id"))
	status := c.FormValue("status")

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid form",
		})
	}

	theradId, err := strconv.Atoi(c.Param("therad_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameters therad_id",
		})
	}

	// if file != null
	if file != nil {
		pathFile := os.Getenv("BASE_URL") + constants.StaticFileTherad + file.Filename
		src, _ := file.Open()
		dir, _ := os.Getwd()
		locationFile := filepath.Join(dir, constants.DirFileTherad, file.Filename)
		dst, _ := os.OpenFile(locationFile, os.O_WRONLY|os.O_CREATE, 06666)
		defer func(dst *os.File) {
			_ = dst.Close()
		}(dst)
		_, _ = io.Copy(dst, src)

		mTherad := model.Therad{
			Judul:            judul,
			Isi:              isi,
			File:             pathFile,
			Status:           status,
			KategoriTheradID: kategoriTheradId,
			UserID:           int(middleware.GetDataFromToken(token)["id"].(float64)),
		}
		therad.UpdateTherad(theradId, mTherad)
	} else {
		mTherad := model.Therad{
			Judul:            judul,
			Isi:              isi,
			Status:           status,
			KategoriTheradID: kategoriTheradId,
			UserID:           int(middleware.GetDataFromToken(token)["id"].(float64)),
		}
		therad.UpdateTherad(theradId, mTherad)
	}

	res := therad.GetTheradById(int(middleware.GetDataFromToken(token)["id"].(float64)), theradId)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Therad pertanyaan anda berhasil di update",
		"data":    res,
	})
}
