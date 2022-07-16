package follow

import (
	"Diskusiaza-BE/app/http/middleware"
	"Diskusiaza-BE/app/repository/follow"
	"Diskusiaza-BE/constants"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetAllUsersController(c echo.Context) error {
	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	res := follow.GetAllUsers(int(middleware.GetDataFromToken(token)["id"].(float64)))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}

func GetUsersByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameters therad_id",
		})
	}
	res := follow.GetUsersById(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}

func GetSearchUsersByNameController(c echo.Context) error {
	key := c.QueryParam("key")
	if key == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Key can't empty",
		})
	}
	res := follow.GetSearchUserByFirstnameOrLastname(key)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"key":  key,
		"data": res,
	})
}
