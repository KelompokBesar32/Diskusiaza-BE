package comment

import (
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/database"
)

func repairResponseComment(mComment []model.CommentResponse) []model.CommentResponse {
	for i := 0; i < len(mComment); i++ {
		mComment[i].AuthorName = getAuthorNameFromComment(mComment[i].ID)
		mComment[i].ReplyCommentDetail = getReplyCommentByCommentId(mComment[i].ID)
	}
	return mComment
}

func getAuthorNameFromComment(commentId uint) string {
	var userResponse model.UserResponse
	database.DB.Raw("SELECT firstname, lastname FROM user "+
		"inner join comment ON comment.user_id = user.id "+
		"WHERE comment.id = ?", commentId).
		Scan(&userResponse)
	authorName := userResponse.Firstname + " " + userResponse.Lastname
	return authorName
}

func CreateComment(mComment model.Comment) {
	database.DB.Create(&mComment)
}

func UpdateComment(commentId int, mComment model.Comment) {
	database.DB.Model(&model.Comment{}).Where("id", commentId).Updates(mComment)
}

func DeleteComment(commentId int) {
	database.DB.Unscoped().Where("id", commentId).Delete(&model.Comment{})
}

func GetCommentByTheradId(theradId int) []model.CommentResponse {
	var mComment []model.CommentResponse
	database.DB.Model(&model.Comment{}).
		Where("therad_id", theradId).
		Scan(&mComment)
	return repairResponseComment(mComment)
}

func GetCommentById(commentId int) model.CommentResponse {
	var mComment model.CommentResponse
	database.DB.Model(&model.Comment{}).
		Where("id", commentId).
		Scan(&mComment)
	mComment.ReplyCommentDetail = getReplyCommentByCommentId(mComment.ID)
	mComment.AuthorName = getAuthorNameFromComment(mComment.ID)
	return mComment
}
