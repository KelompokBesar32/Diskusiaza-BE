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
