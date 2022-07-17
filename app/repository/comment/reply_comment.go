package comment

import (
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/database"
)

func repairResponseReplyComment(mComment []model.ReplyCommentResponse) []model.ReplyCommentResponse {
	for i := 0; i < len(mComment); i++ {
		mComment[i].AuthorName = getAuthorNameFromReplyComment(mComment[i].ID)
	}
	return mComment
}

func getAuthorNameFromReplyComment(replyCommentId uint) string {
	var userResponse model.UserResponse
	database.DB.Raw("SELECT firstname, lastname FROM user "+
		"inner join reply_comment ON "+
		"reply_comment.user_id = user.id WHERE reply_comment.id = ?", replyCommentId).
		Scan(&userResponse)
	authorName := userResponse.Firstname + " " + userResponse.Lastname
	return authorName
}

func getReplyCommentByCommentId(commentId uint) []model.ReplyCommentResponse {
	var mComment []model.ReplyCommentResponse
	database.DB.Model(&model.ReplyComment{}).
		Where("comment_id", commentId).
		Scan(&mComment)
	return repairResponseReplyComment(mComment)
}

func CreateReplyComment(mComment model.ReplyComment) {
	database.DB.Create(&mComment)
}
