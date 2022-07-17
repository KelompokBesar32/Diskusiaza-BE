package ruang

import (
	"Diskusiaza-BE/app/http/middleware"
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/app/repository/ruang"
	"Diskusiaza-BE/constants"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateRuangController(c echo.Context) error {
	mRuang := model.Ruang{}
	_ = c.Bind(&mRuang)

	if mRuang.Judul == "" || mRuang.Deskripsi == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid form input",
		})
	}

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	mRuang.UserID = int(middleware.GetDataFromToken(token)["id"].(float64))

	ruang.CreateRuang(mRuang)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "create ruang successfully",
	})

}

func DeleteRuangController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("ruang_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameters ruang_id",
		})
	}

	ruang.DeleteRuang(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "delete ruang successfully",
	})
}

func UpdateRuangController(c echo.Context) error {
	mRuang := model.Ruang{}
	_ = c.Bind(&mRuang)

	if mRuang.Judul == "" || mRuang.Deskripsi == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "ada inputan yang tidak sesuai",
		})
	}

	ruangID, err := strconv.Atoi(c.Param("ruang_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameters ruang_id",
		})
	}

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	mRuang.UserID = int(middleware.GetDataFromToken(token)["id"].(float64))

	ruang.UpdateRuang(ruangID, mRuang)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update ruang successfully",
	})
}

func GetAllRuangController(c echo.Context) error {
	res := ruang.GetAllRuang()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}

func GetRuangByUserIdController(c echo.Context) error {
	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	userId := int(middleware.GetDataFromToken(token)["id"].(float64))

	res := ruang.GetRuangByUserId(userId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}

func GetRuangByIdController(c echo.Context) error {
	ruangID, err := strconv.Atoi(c.Param("ruang_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameters ruang_id",
		})
	}
	res := ruang.GetRuangById(ruangID)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}

func GetSearchRuangController(c echo.Context) error {
	key := c.QueryParam("key")
	if key == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Key can't empty",
		})
	}
	res := ruang.SearchRuangByName(key)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
		"key":  key,
	})
}

func GetRuangFollowedByUserIdController(c echo.Context) error {
	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	userId := int(middleware.GetDataFromToken(token)["id"].(float64))

	res := ruang.GetRuangFollowedByUserId(userId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}
