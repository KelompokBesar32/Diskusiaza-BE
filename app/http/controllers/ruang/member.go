package ruang

import (
	"Diskusiaza-BE/app/http/middleware"
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/app/repository/ruang"
	"Diskusiaza-BE/constants"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CreateMemberRuangController(c echo.Context) error {
	mMember := model.MemberRuang{}
	_ = c.Bind(&mMember)
	if mMember.RuangID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid form input",
		})
	}

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	userId := int(middleware.GetDataFromToken(token)["id"].(float64))
	mMember.UserID = userId

	if ruang.CheckIfUserAuthorInRuang(mMember.RuangID, userId) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "your can't join in this ruang, because your is author",
		})
	}

	if !ruang.CheckIfUserJoinInRuang(mMember) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "your already join in ruang",
		})
	}

	ruang.CreateMemberRuang(mMember)
	ruangDetail := ruang.GetRuangById(mMember.RuangID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "your successfully join in this ruang",
		"ruang_detail": ruangDetail,
	})
}

func DeleteMemberRuangController(c echo.Context) error {
	mMember := model.MemberRuang{}
	ruangId, err := strconv.Atoi(c.Param("ruang_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameters ruang_id",
		})
	}
	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	userId := int(middleware.GetDataFromToken(token)["id"].(float64))

	mMember.UserID = userId
	mMember.RuangID = ruangId
	ruang.DeleteMemberRuang(mMember)
	ruangDetail := ruang.GetRuangById(mMember.RuangID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "your successfully leave from this ruang",
		"ruang_detail": ruangDetail,
	})
}

func GetAllMemberInRuangController(c echo.Context) error {
	ruangId, err := strconv.Atoi(c.Param("ruang_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameters ruang_id",
		})
	}
	res := ruang.GetAllMemberInRuang(ruangId)
	ruangDetail := ruang.GetRuangById(ruangId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"member":       res,
		"ruang_detail": ruangDetail,
	})
}
