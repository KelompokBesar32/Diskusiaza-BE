package controllers

import (
	"Diskusiaza-BE/app/http/middleware"
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/app/repository"
	"Diskusiaza-BE/constants"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterController(c echo.Context) error {
	validate := validator.New()
	user := model.User{}
	_ = c.Bind(&user)

	if err := validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "ada inputan yang tidak sesuai",
		})
	}
	if !repository.GetUsersByEmail(user.Email) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "email " + user.Email + " telah digunakan",
		})
	}

	repository.CreateUsers(user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "proses register berhasil",
	})
}

func LoginController(c echo.Context) error {
	validate := validator.New()
	user := model.User{}
	_ = c.Bind(&user)

	if err := validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "ada inputan yang tidak sesuai",
		})
	}

	mUser, err := repository.Login(user)

	if !err {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "email atau password salah",
		})
	}

	token, _ := middleware.CreateTokenUsers(mUser)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func LogoutController(c echo.Context) error {
	// invalidate token ??
	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	repository.RemoveToken(token)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user successfully logout",
	})
}
