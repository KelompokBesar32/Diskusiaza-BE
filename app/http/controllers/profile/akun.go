package profile

import (
	"Diskusiaza-BE/app/http/middleware"
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/app/repository/profile"
	"Diskusiaza-BE/constants"
	"Diskusiaza-BE/library"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func getDataFromToken(token string) map[string]interface{} {
	claimsToken, _ := middleware.DecodeTokenUsers(token)
	dataFromToken := claimsToken["data"].(map[string]interface{})
	return dataFromToken
}

func GetUsersByToken(c echo.Context) error {
	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	res := profile.GetDetailUser(int(getDataFromToken(token)["id"].(float64)))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}

func UpdateUsersData(c echo.Context) error {
	firstName := c.FormValue("firstname")
	lastName := c.FormValue("lastname")
	nohp := c.FormValue("nohp")
	tglLahir := c.FormValue("tanggal_lahir")
	jenisKelamin := c.FormValue("jenis_kelamin")
	foto, _ := c.FormFile("foto")

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]

	// if foto != null
	if foto != nil {
		pathImage := os.Getenv("BASE_URL") + constants.StaticFileUsersFoto + foto.Filename
		if !library.CheckExtensionForImage(filepath.Ext(pathImage)) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "invalid extension file",
			})
		}
		src, _ := foto.Open()
		dir, _ := os.Getwd()
		locationFile := filepath.Join(dir, constants.DirFileUsersFoto, foto.Filename)
		dst, _ := os.OpenFile(locationFile, os.O_WRONLY|os.O_CREATE, 06666)
		defer dst.Close()
		io.Copy(dst, src)

		mUser := model.User{
			Firstname:    firstName,
			Lastname:     lastName,
			Nohp:         nohp,
			Foto:         pathImage,
			TanggalLahir: tglLahir,
			JenisKelamin: jenisKelamin,
		}
		profile.UpdateUser(int(getDataFromToken(token)["id"].(float64)), mUser)

	} else {
		mUser := model.User{
			Firstname:    firstName,
			Lastname:     lastName,
			Nohp:         nohp,
			TanggalLahir: tglLahir,
			JenisKelamin: jenisKelamin,
		}
		profile.UpdateUser(int(getDataFromToken(token)["id"].(float64)), mUser)
	}

	res := profile.GetDetailUser(int(getDataFromToken(token)["id"].(float64)))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data users successfully updated",
		"data":    res,
	})
}
