package follow

import (
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/database"
)

func GetDetailFollowed(userId int) []model.UserResponse {
	var mUser []model.UserResponse
	database.DB.Model(&model.User{}).
		Select("user.id", "firstname", "lastname", "email", "nohp", "foto", "tanggal_lahir", "jenis_kelamin", "followers_id", "followed_id").
		Joins("inner join followers ON followers.followed_id = user.id").
		Where("followers.followers_id", userId).
		Scan(&mUser)
	return repairResponse(mUser)
}

func GetDetailFollowers(userId int) []model.UserResponse {
	var mUser []model.UserResponse
	database.DB.Model(&model.User{}).
		Select("user.id", "firstname", "lastname", "email", "nohp", "foto", "tanggal_lahir", "jenis_kelamin", "followers_id", "followed_id").
		Joins("inner join followers ON followers.followers_id = user.id").
		Where("followers.followed_id", userId).
		Scan(&mUser)
	return repairResponse(mUser)
}

func CheckIsFollowing(mFollowers model.Followers) bool {
	var total int
	database.DB.Raw("SELECT COUNT(followers_id) FROM followers WHERE followers_id = ? AND followed_id = ?", mFollowers.FollowersID, mFollowers.FollowedID).
		Scan(&total)
	if total == 0 {
		return true
	}
	return false
}

func ActionFollowUsers(mFollowers model.Followers) {
	database.DB.Create(&mFollowers)
}

func ActionUnFollowUsers(mFollowers model.Followers) {
	database.DB.Unscoped().
		Where("followers_id", mFollowers.FollowersID).
		Where("followed_id", mFollowers.FollowedID).
		Delete(&model.Followers{})
}
