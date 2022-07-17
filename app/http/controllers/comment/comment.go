package comment

import (
	"Diskusiaza-BE/app/http/middleware"
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/app/repository/comment"
	therad2 "Diskusiaza-BE/app/repository/therad"
	"Diskusiaza-BE/constants"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func CreateCommentController(c echo.Context) error {
	isi := c.FormValue("isi")
	theradId, err := strconv.Atoi(c.FormValue("therad_id"))
	file, _ := c.FormFile("file")

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid form",
		})
	}

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]

	if file != nil {
		pathFile := os.Getenv("BASE_URL") + constants.StaticFileComment + file.Filename
		src, _ := file.Open()
		dir, _ := os.Getwd()
		locationFile := filepath.Join(dir, constants.DirFileComment, file.Filename)
		dst, _ := os.OpenFile(locationFile, os.O_WRONLY|os.O_CREATE, 06666)
		defer func(dst *os.File) {
			_ = dst.Close()
		}(dst)
		_, _ = io.Copy(dst, src)

		mComment := model.Comment{
			Isi:      isi,
			File:     pathFile,
			TheradID: theradId,
			UserID:   int(middleware.GetDataFromToken(token)["id"].(float64)),
		}

		comment.CreateComment(mComment)

	} else {

		mComment := model.Comment{
			Isi:      isi,
			TheradID: theradId,
			UserID:   int(middleware.GetDataFromToken(token)["id"].(float64)),
		}

		comment.CreateComment(mComment)

	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "your comment successfully published in this therad",
	})

}

func UpdateCommentController(c echo.Context) error {
	isi := c.FormValue("isi")
	theradId, errTheradId := strconv.Atoi(c.FormValue("therad_id"))
	file, _ := c.FormFile("file")
	commentId, errComment := strconv.Atoi(c.Param("comment_id"))

	if errTheradId != nil || errComment != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid form",
		})
	}

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]

	if file != nil {
		pathFile := os.Getenv("BASE_URL") + constants.StaticFileComment + file.Filename
		src, _ := file.Open()
		dir, _ := os.Getwd()
		locationFile := filepath.Join(dir, constants.DirFileComment, file.Filename)
		dst, _ := os.OpenFile(locationFile, os.O_WRONLY|os.O_CREATE, 06666)
		defer func(dst *os.File) {
			_ = dst.Close()
		}(dst)
		_, _ = io.Copy(dst, src)

		mComment := model.Comment{
			Isi:      isi,
			File:     pathFile,
			TheradID: theradId,
			UserID:   int(middleware.GetDataFromToken(token)["id"].(float64)),
		}

		comment.UpdateComment(commentId, mComment)

	} else {

		mComment := model.Comment{
			Isi:      isi,
			TheradID: theradId,
			UserID:   int(middleware.GetDataFromToken(token)["id"].(float64)),
		}

		comment.UpdateComment(commentId, mComment)

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "your comment successfully updated in this therad",
	})
}

func DeleteCommentController(c echo.Context) error {
	commentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameter comment_id",
		})
	}
	comment.DeleteComment(commentId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "your comment successfully deleted",
	})
}

func GetCommentByTheradIdController(c echo.Context) error {
	theradId, err := strconv.Atoi(c.Param("therad_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameter therad_id",
		})
	}

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TokenJwtType):]
	userId := int(middleware.GetDataFromToken(token)["id"].(float64))

	res := comment.GetCommentByTheradId(theradId)
	therad := therad2.GetTheradById(userId, theradId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"comment":       res,
		"therad_detail": therad,
	})
}

func GetCommentByIdController(c echo.Context) error {
	commentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameter comment_id",
		})
	}

	res := comment.GetCommentById(commentId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"comment": res,
	})
}
