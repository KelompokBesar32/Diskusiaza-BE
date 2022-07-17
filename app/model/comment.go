package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	Isi      string `json:"isi"`
	File     string `json:"file"`
	UserID   int    `json:"user_id"`
	TheradID int    `json:"therad_id"`
	User     User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Therad   Therad `gorm:"foreignKey:TheradID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (Comment) TableName() string {
	return "comment"
}

type CommentResponse struct {
	ID                 uint                   `gorm:"primaryKey" json:"id"`
	Isi                string                 `json:"isi"`
	File               string                 `json:"file"`
	AuthorName         string                 `json:"author_name"`
	ReplyCommentDetail []ReplyCommentResponse `json:"reply_comment_detail"`
}
