package model

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	Isi      string `json:"isi"`
	TheradID int    `json:"therad_id"`
	UserID   int    `json:"user_id"`
	Therad   Therad `gorm:"foreignKey:TheradID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	User     User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (Report) TableName() string {
	return "report"
}
