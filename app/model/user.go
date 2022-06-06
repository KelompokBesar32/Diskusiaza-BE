package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Nohp      string `json:"nohp"`
	Foto      string `json:"foto"`
	Token     string `json:"token"`
	Password  string `json:"password"`
	RoleID    int    `json:"role_id"`
	Role      Role   `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (User) TableName() string {
	return "user"
}
