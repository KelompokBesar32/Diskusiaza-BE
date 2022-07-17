package model

import "gorm.io/gorm"

type Therad struct {
	gorm.Model
	ID               uint           `gorm:"primaryKey" json:"id"`
	Judul            string         `json:"judul"`
	Isi              string         `json:"isi"`
	File             string         `json:"file"`
	Dilihat          int            `json:"dilihat"`
	Status           string         `json:"status" gorm:"type:enum('active', 'not_active')"`
	UserID           int            `json:"user_id"`
	KategoriTheradID int            `json:"kategori_therad_id"`
	RuangID          int            `json:"ruang_id"`
	User             User           `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	KategoriTherad   KategoriTherad `gorm:"foreignKey:KategoriTheradID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Ruang            Ruang          `gorm:"foreignKey:RuangID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (Therad) TableName() string {
	return "therad"
}

type TheradResponse struct {
	ID           uint   `json:"id"`
	UserId       int    `json:"user_id"`
	Judul        string `json:"judul"`
	Isi          string `json:"isi"`
	File         string `json:"file"`
	Dilihat      int    `json:"dilihat"`
	Status       string `json:"status"`
	KategoriName string `json:"kategori_name"`
	AuthorName   string `json:"author_name" gorm:"->;type:GENERATED ALWAYS AS (concat(firstname,' ',lastname));default:(-);"` // firstname + lastname
	RuangName    string `json:"ruang_name"`
	TotalLike    int    `json:"total_like"`
	IsLike       bool   `json:"is_like"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
