package therad

import (
	"Diskusiaza-BE/app/http/middleware"
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/app/repository/therad"
	"Diskusiaza-BE/constants"
	"github.com/labstack/echo/v4"
	"net/http"
)

func LikeTheradController(c echo.Context) error {
	mLike := model.Like{}
	_ = c.Bind(&mLike)
	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	mLike.UserID = int(middleware.GetDataFromToken(token)["id"].(float64))

	if therad.CheckIfUserLikeTherad(int(middleware.GetDataFromToken(token)["id"].(float64)), mLike.TheradID) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "sorry, your already like this therad",
			"data":    therad.GetTheradById(mLike.UserID, int(mLike.TheradID)),
		})
	}
	therad.CreateLikeTherad(mLike)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "anda berhasil menyukai therad ini",
		"data":    therad.GetTheradById(mLike.UserID, int(mLike.TheradID)),
	})
}

func UnLikeTheradController(c echo.Context) error {
	mLike := model.Like{}
	_ = c.Bind(&mLike)
	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	therad.DeleteLikeTherad(int(mLike.TheradID), int(middleware.GetDataFromToken(token)["id"].(float64)))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "anda berhasil unlike therad ini",
		"data":    therad.GetTheradById(int(middleware.GetDataFromToken(token)["id"].(float64)), int(mLike.TheradID)),
	})
}
