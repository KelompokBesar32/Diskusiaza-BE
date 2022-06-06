package model

import "gorm.io/gorm"

type MemberRuang struct {
	gorm.Model
	ID      uint  `gorm:"primaryKey" json:"id"`
	RuangID int   `json:"ruang_id"`
	UserID  int   `json:"user_id"`
	Ruang   Ruang `gorm:"foreignKey:RuangID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	User    User  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (MemberRuang) TableName() string {
	return "member_ruang"
}
