package follow

import (
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/database"
)

func repairResponse(mUser []model.UserResponse) []model.UserResponse {
	for i := 0; i < len(mUser); i++ {
		mUser[i].TotalFollowed = getTotalFollowed(mUser[i].ID)
		mUser[i].TotalFollowers = getTotalFollowers(mUser[i].ID)
		mUser[i].TotalTheradCreated = getTotalTheradCreated(mUser[i].ID)
	}
	return mUser
}

func getTotalTheradCreated(userId int) int {
	var totalTherad int
	database.DB.Raw("SELECT COUNT(user_id) FROM therad WHERE user_id = ?", userId).Scan(&totalTherad)
	return totalTherad
}

func getTotalFollowers(userId int) int {
	var totalFollowers int
	database.DB.Raw("SELECT COUNT(followers_id) FROM followers WHERE followed_id = ?", userId).Scan(&totalFollowers)
	return totalFollowers
}

func getTotalFollowed(userId int) int {
	var totalFollowed int
	database.DB.Raw("SELECT COUNT(followed_id) FROM followers WHERE followers_id = ?", userId).Scan(&totalFollowed)
	return totalFollowed
}

func GetAllUsers(userId int) []model.UserResponse {
	var mUser []model.UserResponse
	database.DB.Raw("SELECT *FROM user WHERE id != ? ORDER BY firstname ASC", userId).Scan(&mUser)
	return repairResponse(mUser)
}

func GetSearchUserByFirstnameOrLastname(key string) []model.UserResponse {
	var mUser []model.UserResponse
	database.DB.Table("user").Where("firstname LIKE ? OR lastname LIKE ?", "%"+key+"%", "%"+key+"%").Scan(&mUser)
	return repairResponse(mUser)
}

func GetUsersById(userId int) model.UserResponse {
	var mUser model.UserResponse
	database.DB.First(&model.User{}, "id = ?", userId).Scan(&mUser)
	mUser.TotalFollowers = getTotalFollowers(userId)
	mUser.TotalFollowed = getTotalFollowed(userId)
	mUser.TotalTheradCreated = getTotalTheradCreated(userId)
	return mUser
}
