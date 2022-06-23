package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	TheradID uint   `json:"therad_id"`
	UserID   int    `json:"user_id"`
	Therad   Therad `gorm:"foreignKey:TheradID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	User     User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (Like) TableName() string {
	return "likes"
}
