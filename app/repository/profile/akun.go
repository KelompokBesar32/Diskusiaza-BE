package profile

import (
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/database"
)

func getTotalFollowers(userId int) int {
	var totalFollowers int
	database.DB.Raw("SELECT COUNT(followed_id) FROM followers WHERE followed_id = ?", userId).Scan(&totalFollowers)
	return totalFollowers
}

func getTotalFollowed(userId int) int {
	var totalFollowed int
	database.DB.Raw("SELECT COUNT(followers_id) FROM followers WHERE followers_id = ?", userId).Scan(&totalFollowed)
	return totalFollowed
}

func getTotalTheradCreated(userId int) int {
	var totalTherad int
	database.DB.Raw("SELECT COUNT(user_id) FROM therad WHERE user_id = ?", userId).Scan(&totalTherad)
	return totalTherad
}

func GetDetailUser(userId int) model.UserResponse {
	var mUser model.UserResponse
	database.DB.First(&model.User{}, "id = ?", userId).Scan(&mUser)
	mUser.TotalFollowers = getTotalFollowers(userId)
	mUser.TotalFollowed = getTotalFollowed(userId)
	mUser.TotalTheradCreated = getTotalTheradCreated(userId)
	return mUser
}

func UpdateUser(id int, mUser model.User) {
	database.DB.Model(&model.User{}).Where("id", id).Updates(mUser)
}
