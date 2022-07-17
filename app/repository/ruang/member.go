package ruang

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

func CreateMemberRuang(mMember model.MemberRuang) {
	database.DB.Create(&mMember)
}

func DeleteMemberRuang(mMember model.MemberRuang) {
	database.DB.Unscoped().Where("user_id", mMember.UserID).
		Where("ruang_id", mMember.RuangID).
		Delete(&model.MemberRuang{})
}

func CheckIfUserJoinInRuang(mMember model.MemberRuang) bool {
	var total int
	database.DB.Raw("SELECT COUNT(id) FROM member_ruang WHERE ruang_id = ? AND user_id = ?", mMember.RuangID, mMember.UserID).
		Scan(&total)
	if total == 1 {
		return false
	}
	return true
}

func CheckIfUserAuthorInRuang(ruangId, userId int) bool {
	var res int
	database.DB.Raw("SELECT user_id FROM ruang WHERE id = ?", ruangId).Scan(&res)
	if res == userId {
		return true
	}
	return false
}

func GetAllMemberInRuang(ruangId int) []model.UserResponse {
	var mUser []model.UserResponse
	database.DB.Model(&model.User{}).
		Select("user.id", "firstname", "lastname", "email", "nohp", "foto", "tanggal_lahir", "jenis_kelamin").
		Joins("inner join member_ruang ON member_ruang.user_id = user.id").
		Where("member_ruang.ruang_id", ruangId).
		Scan(&mUser)
	return repairResponse(mUser)
}
