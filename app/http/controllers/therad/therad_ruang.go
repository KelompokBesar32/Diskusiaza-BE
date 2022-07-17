package therad

import (
	"Diskusiaza-BE/app/http/middleware"
	"Diskusiaza-BE/app/model"
	ruang2 "Diskusiaza-BE/app/repository/ruang"
	"Diskusiaza-BE/app/repository/therad"
	"Diskusiaza-BE/constants"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func CreateTheradInRuangController(c echo.Context) error {
	judul := c.FormValue("judul")
	isi := c.FormValue("isi")
	file, _ := c.FormFile("file")
	kategoriTheradId, errKategori := strconv.Atoi(c.FormValue("kategori_therad_id"))
	ruangId, errRuang := strconv.Atoi(c.FormValue("ruang_id"))

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]

	if errKategori != nil || errRuang != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid form",
		})
	}

	ruang := ruang2.GetRuangById(ruangId)
	if ruang.ID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "ruang not found",
		})
	}
	mMember := model.MemberRuang{
		RuangID: int(ruang.ID),
		UserID:  int(middleware.GetDataFromToken(token)["id"].(float64)),
	}
	if ruang2.CheckIfUserJoinInRuang(mMember) && !ruang2.CheckIfUserAuthorInRuang(mMember.RuangID, mMember.UserID) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "can't create therad in this ruang, because your not member in this ruang",
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
			RuangID:          ruangId,
			UserID:           int(middleware.GetDataFromToken(token)["id"].(float64)),
		}
		therad.CreateTheradInRuang(mTherad)
	} else {
		mTherad := model.Therad{
			Judul:            judul,
			Isi:              isi,
			Dilihat:          0,
			Status:           "active",
			KategoriTheradID: kategoriTheradId,
			RuangID:          ruangId,
			UserID:           int(middleware.GetDataFromToken(token)["id"].(float64)),
		}
		therad.CreateTheradInRuang(mTherad)
	}

	mRuang := ruang2.GetRuangById(ruangId)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":      "Therad pertanyaan anda berhasil diposting di ruang ini",
		"ruang_detail": mRuang,
	})
}

func GetAllTheradInRuangController(c echo.Context) error {
	ruangId, err := strconv.Atoi(c.Param("ruang_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameters ruang_id",
		})
	}

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	userId := int(middleware.GetDataFromToken(token)["id"].(float64))

	res := therad.GetAllListTheradInRuang(userId, ruangId)
	mRuang := ruang2.GetRuangById(ruangId)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"therad":       res,
		"ruang_detail": mRuang,
	})

}
