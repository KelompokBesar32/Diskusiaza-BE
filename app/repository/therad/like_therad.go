package therad

import (
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/database"
)

func CheckIfUserLikeTherad(userId int, theradId uint) bool {
	var mLike model.Like
	database.DB.Model(&model.Like{}).Where("user_id", userId).Where("therad_id", theradId).Scan(&mLike)
	if mLike.UserID == 0 {
		return false
	}
	return true
}

func CreateLikeTherad(mLike model.Like) {
	database.DB.Create(&mLike)
}

func DeleteLikeTherad(theradId, userId int) {
	database.DB.Unscoped().Where("therad_id", theradId).Where("user_id", userId).Delete(&model.Like{})
}
