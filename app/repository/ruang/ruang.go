package ruang

import (
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/database"
)

func repairRuangResponse(mRuang []model.RuangResponse) []model.RuangResponse {
	for i := 0; i < len(mRuang); i++ {
		mRuang[i].TotalTheradCreated = totalTheradCreated(mRuang[i].ID)
		mRuang[i].TotalMember = totalMember(mRuang[i].ID)
		mRuang[i].AuthorName = getAuthorNameFromRuang(mRuang[i].ID)
	}
	return mRuang
}

func totalTheradCreated(ruangId uint) int {
	var total int
	database.DB.Raw("SELECT COUNT(id) FROM therad WHERE ruang_id = ?", ruangId).Scan(&total)
	return total
}

func totalMember(ruangId uint) int {
	var total int
	database.DB.Raw("SELECT COUNT(id) FROM member_ruang WHERE ruang_id = ?", ruangId).Scan(&total)
	return total
}

func getAuthorNameFromRuang(ruangId uint) string {
	var userResponse model.UserResponse
	database.DB.Raw("SELECT firstname, lastname FROM user inner join ruang ON ruang.user_id = user.id WHERE ruang.id = ?", ruangId).
		Scan(&userResponse)
	authorName := userResponse.Firstname + " " + userResponse.Lastname
	return authorName
}

func CreateRuang(mRuang model.Ruang) {
	database.DB.Create(&mRuang)
}

func DeleteRuang(ruangId int) {
	database.DB.Unscoped().Where("id", ruangId).Delete(&model.Ruang{})
}

func UpdateRuang(ruangId int, mRuang model.Ruang) {
	database.DB.Model(&model.Ruang{}).Where("id", ruangId).Updates(mRuang)
}

func GetAllRuang() []model.RuangResponse {
	var mRuang []model.RuangResponse
	database.DB.Raw("SELECT *FROM ruang ORDER BY id DESC").Scan(&mRuang)
	return repairRuangResponse(mRuang)
}

func GetRuangByUserId(userId int) []model.RuangResponse {
	var mRuang []model.RuangResponse
	database.DB.Raw("SELECT *FROM ruang WHERE user_id = ? ORDER BY id DESC", userId).Scan(&mRuang)
	return repairRuangResponse(mRuang)
}

func GetRuangById(ruangId int) model.RuangResponse {
	var mRuang model.RuangResponse
	database.DB.Model(&model.Ruang{}).Where("id", ruangId).Scan(&mRuang)
	mRuang.TotalMember = totalMember(mRuang.ID)
	mRuang.TotalTheradCreated = totalTheradCreated(mRuang.ID)
	mRuang.AuthorName = getAuthorNameFromRuang(mRuang.ID)
	return mRuang
}

func SearchRuangByName(key string) []model.RuangResponse {
	var mRuang []model.RuangResponse
	database.DB.Table("ruang").Where("judul LIKE ?", "%"+key+"%").Scan(&mRuang)
	return repairRuangResponse(mRuang)
}
