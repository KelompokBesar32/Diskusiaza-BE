package model

import "gorm.io/gorm"

type KategoriTherad struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey" json:"id"`
	KategoriName string `json:"kategori_name"`
}

func (KategoriTherad) TableName() string {
	return "kategori_therad"
}

type KategoriTheradResponse struct {
	ID           uint   `json:"id"`
	KategoriName string `json:"kategori_name"`
}
