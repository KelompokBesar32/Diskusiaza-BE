package model

import "gorm.io/gorm"

type ReplyComment struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey" json:"id"`
	Isi       string  `json:"isi"`
	File      string  `json:"file"`
	UserID    int     `json:"user_id"`
	CommentID int     `json:"comment_id"`
	User      User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Comment   Comment `gorm:"foreignKey:CommentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (ReplyComment) TableName() string {
	return "reply_comment"
}
