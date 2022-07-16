package follow

import (
	"Diskusiaza-BE/app/http/middleware"
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/app/repository/follow"
	"Diskusiaza-BE/constants"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetDetailFollowedController(c echo.Context) error {
	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	res := follow.GetDetailFollowed(int(middleware.GetDataFromToken(token)["id"].(float64)))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "detail your followed",
		"data":    res,
	})
}

func GetDetailFollowersController(c echo.Context) error {
	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	res := follow.GetDetailFollowers(int(middleware.GetDataFromToken(token)["id"].(float64)))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "detail your followers",
		"data":    res,
	})
}

func FollowedUserController(c echo.Context) error {
	mFollowers := model.Followers{}
	_ = c.Bind(&mFollowers)

	if mFollowers.FollowedID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "followed id can't be null",
		})
	}

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	userId := int(middleware.GetDataFromToken(token)["id"].(float64))

	mFollowers.FollowersID = userId

	if !follow.CheckIsFollowing(mFollowers) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "your already follow this user",
		})
	}

	follow.ActionFollowUsers(mFollowers)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "your successfully follow this user",
	})
}

func UnFollowedUserController(c echo.Context) error {
	mFollowers := model.Followers{}
	c.Bind(&mFollowers)

	if mFollowers.FollowedID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "followed id can't be null",
		})
	}

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	userId := int(middleware.GetDataFromToken(token)["id"].(float64))

	mFollowers.FollowersID = userId

	if follow.CheckIsFollowing(mFollowers) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "your already unfollow this user",
		})
	}

	follow.ActionUnFollowUsers(mFollowers)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "unfollow user successfully",
	})
}

func RemoveFollowersController(c echo.Context) error {
	mFollowers := model.Followers{}
	c.Bind(&mFollowers)

	if mFollowers.FollowersID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "followers id can't be null",
		})
	}

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	userId := int(middleware.GetDataFromToken(token)["id"].(float64))

	mFollowers.FollowedID = userId

	if follow.CheckIsFollowing(mFollowers) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "your already remove this user from your followers",
		})
	}

	follow.ActionUnFollowUsers(mFollowers)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "remove followers successfully",
	})
}
