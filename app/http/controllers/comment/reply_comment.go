package comment

import (
	"Diskusiaza-BE/app/http/middleware"
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/app/repository/comment"
	"Diskusiaza-BE/constants"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func CreateReplyCommentController(c echo.Context) error {
	isi := c.FormValue("isi")
	commentId, err := strconv.Atoi(c.FormValue("comment_id"))
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

		mComment := model.ReplyComment{
			Isi:       isi,
			File:      pathFile,
			CommentID: commentId,
			UserID:    int(middleware.GetDataFromToken(token)["id"].(float64)),
		}

		comment.CreateReplyComment(mComment)

	} else {

		mComment := model.ReplyComment{
			Isi:       isi,
			CommentID: commentId,
			UserID:    int(middleware.GetDataFromToken(token)["id"].(float64)),
		}

		comment.CreateReplyComment(mComment)

	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "reply this comment successfully",
	})
}
