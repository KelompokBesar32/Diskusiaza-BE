package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//Comment struct
type Comment struct {
	ID               uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" groups:"comment,orgComments"`
	Value            string         `gorm:"varchar(1000)" json:"value" groups:"comment,orgComments"`
	UserID           int            `json:"user_id"`
	User             User           `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	KategoriTheradID int            `json:"kategori_therad_id"`
	RuangID          int            `json:"ruang_id"`
	CreatedAt        time.Time      `gorm:"" json:"created_at" groups:"comment,orgComments"`
	UpdatedAt        time.Time      `gorm:"" json:"updated_at" groups:"comment,orgComments"`
	DeletedAt        *time.Time     `gorm:"" json:"deleted_at" groups:""`
	KategoriTherad   KategoriTherad `gorm:"foreignKey:KategoriTheradID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Ruang            Ruang          `gorm:"foreignKey:RuangID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

//CommentCreate struct
type CommentCreate struct {
	Value string `valid:"stringlength(3|1000)~COMMENT_VALUE_INVALID" json:"value" groups:"comment"`
}

//CommentsPagination struct
type CommentsPagination struct {
	Comments []Comment `json:"comments" groups:"orgComments"`
}
