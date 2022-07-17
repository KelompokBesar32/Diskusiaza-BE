package model

import "gorm.io/gorm"

type Ruang struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
	UserID    int    `json:"user_id"`
	User      User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (Ruang) TableName() string {
	return "ruang"
}

type RuangResponse struct {
	ID                 uint   `gorm:"primaryKey" json:"id"`
	Judul              string `json:"judul"`
	Deskripsi          string `json:"deskripsi"`
	AuthorName         string `json:"author_name"`
	TotalTheradCreated int    `json:"total_therad_created"`
	TotalMember        int    `json:"total_member"`
}
