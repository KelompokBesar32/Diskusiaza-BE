package model

import "gorm.io/gorm"

type Followers struct {
	gorm.Model
	ID              uint `gorm:"primaryKey" json:"id"`
	FollowersID     int  `json:"followers_id"`
	FollowedID      int  `json:"followed_id"`
	UserFollowersID User `gorm:"foreignKey:FollowersID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	UserFollowedID  User `gorm:"foreignKey:FollowedID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (Followers) TableName() string {
	return "followers"
}
