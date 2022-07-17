package therad

//type CommentCreateResponse struct {
//	Message string        `json:"message"`
//	Comment model.Comment `json:"comment"`
//}
//
////CommentsGetResponse struct
//type CommentsGetResponse struct {
//	Message string `json:"message" groups:"orgComments"`
//	model.CommentsPagination
//}
//
////CommentsDeleteResponse struct
//type CommentsDeleteResponse struct {
//	Message     string `json:"message" groups:"deleteComments"`
//	DeleteCount int64  `json:"delete_count" groups:"deleteComments"`
//}
//
//func CreateComment(c echo.Context) error {
//	validate := validator.New()
//	comment := model.Comment{}
//	_ = c.Bind(&comment)
//
//	if err := validate.Struct(comment); err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]interface{}{
//			"message": "ada inputan yang tidak sesuai",
//		})
//	}
//	userID := c.Get("user_id").(int)
//	comment.UserID = userID
//	comment.CreatedAt = time.Now()
//	comment.UpdatedAt = time.Now()
//	database.DB.Create(&comment)
//
//	return c.JSON(http.StatusOK, map[string]interface{}{
//		"message": "proses create comment berhasil",
//		"comment": comment,
//	})
//}
//
//func updateComment(c echo.Context) error {
//	validate := validator.New()
//	comment := model.Comment{}
//	_ = c.Bind(&comment)
//
//	if err := validate.Struct(comment); err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]interface{}{
//			"message": "ada inputan yang tidak sesuai",
//		})
//	}
//
//	database.DB.Save(&comment)
//
//	return c.JSON(http.StatusOK, map[string]interface{}{
//		"message": "proses update comment berhasil",
//	})
//}
//
//func deleteComment(c echo.Context) error {
//	comment := model.Comment{}
//	_ = c.Bind(&comment)
//
//	database.DB.Delete(&comment)
//
//	return c.JSON(http.StatusOK, map[string]interface{}{
//		"message": "proses delete comment berhasil",
//	})
//}
//
//func getComments(c echo.Context) error {
//	comment := model.Comment{}
//	_ = c.Bind(&comment)
//
//	database.DB.Find(&comment)
//
//	return c.JSON(http.StatusOK, map[string]interface{}{
//		"message": "proses get comment berhasil",
//		"comment": comment,
//	})
//}
