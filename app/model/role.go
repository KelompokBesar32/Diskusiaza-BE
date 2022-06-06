package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	RoleName string `json:"role_name"`
}

func (Role) TableName() string {
	return "role"
}
