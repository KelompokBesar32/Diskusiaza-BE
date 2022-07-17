package therad

import (
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/database"
)

func repairTheradResponse(userId int, theradResponse []model.TheradResponse) []model.TheradResponse {
	for i := 0; i < len(theradResponse); i++ {
		if CheckIfUserLikeTherad(userId, theradResponse[i].ID) {
			theradResponse[i].IsLike = true
		}
	}
	return theradResponse
}

func CreateNormalTherad(therad model.Therad) {
	database.DB.Omit("ruang_id").Create(&therad)
}

func CreateTheradInRuang(therad model.Therad) {
	database.DB.Create(&therad)
}

func GetListTheradByUserId(userId int) []model.TheradResponse {
	var theradResponse []model.TheradResponse
	database.DB.Model(&model.Therad{}).
		Select("therad.id", "therad.judul", "therad.user_id", "therad.isi", "therad.file", "therad.dilihat",
			"therad.status", "kategori_therad.kategori_name", "ruang.judul AS ruang_name",
			"CONCAT(user.firstname, ' ' ,user.lastname) AS author_name",
			"COUNT(likes.therad_id) AS total_like",
			"therad.created_at", "therad.updated_at").
		Joins("inner join kategori_therad ON kategori_therad.id = therad.kategori_therad_id").
		Joins("inner join user ON user.id = therad.user_id").
		Joins("left join ruang ON ruang.id = therad.ruang_id").
		Joins("left join likes ON likes.therad_id = therad.id").
		Where("therad.user_id", userId).
		Group("therad.id").
		Scan(&theradResponse)
	return repairTheradResponse(userId, theradResponse)
}

func GetTheradById(userId, theradId int) model.TheradResponse {
	var theradResponse model.TheradResponse
	database.DB.Model(&model.Therad{}).
		Select("therad.id", "therad.judul", "therad.user_id", "therad.isi", "therad.file", "therad.dilihat",
			"therad.status", "kategori_therad.kategori_name", "ruang.judul AS ruang_name",
			"CONCAT(user.firstname, ' ' ,user.lastname) AS author_name",
			"COUNT(likes.therad_id) AS total_like",
			"therad.created_at", "therad.updated_at").
		Joins("inner join kategori_therad ON kategori_therad.id = therad.kategori_therad_id").
		Joins("inner join user ON user.id = therad.user_id").
		Joins("left join ruang ON ruang.id = therad.ruang_id").
		Joins("left join likes ON likes.therad_id = therad.id").
		Where("therad.id", theradId).
		Group("therad.id").
		Scan(&theradResponse)
	UpdateTherad(theradId, model.Therad{Dilihat: theradResponse.Dilihat + 1}) // +1
	if CheckIfUserLikeTherad(userId, theradResponse.ID) {
		theradResponse.IsLike = true
	}
	return theradResponse
}

func GetAllListTherad(userId int) []model.TheradResponse {
	var theradResponse []model.TheradResponse
	database.DB.Model(&model.Therad{}).
		Select("therad.id", "therad.judul", "therad.user_id", "therad.isi", "therad.file", "therad.dilihat",
			"therad.status", "kategori_therad.kategori_name", "ruang.judul AS ruang_name",
			"CONCAT(user.firstname, ' ' ,user.lastname) AS author_name",
			"COUNT(likes.therad_id) AS total_like",
			"therad.created_at", "therad.updated_at").
		Joins("inner join kategori_therad ON kategori_therad.id = therad.kategori_therad_id").
		Joins("inner join user ON user.id = therad.user_id").
		Joins("left join ruang ON ruang.id = therad.ruang_id").
		Joins("left join likes ON likes.therad_id = therad.id").
		Where("status", "active").
		Group("therad.id").
		Scan(&theradResponse)
	return repairTheradResponse(userId, theradResponse)
}

func GetAllListTheradInRuang(userId, ruangId int) []model.TheradResponse {
	var theradResponse []model.TheradResponse
	database.DB.Model(&model.Therad{}).
		Select("therad.id", "therad.judul", "therad.user_id", "therad.isi", "therad.file", "therad.dilihat",
			"therad.status", "kategori_therad.kategori_name", "ruang.judul AS ruang_name",
			"CONCAT(user.firstname, ' ' ,user.lastname) AS author_name",
			"COUNT(likes.therad_id) AS total_like",
			"therad.created_at", "therad.updated_at").
		Joins("inner join kategori_therad ON kategori_therad.id = therad.kategori_therad_id").
		Joins("inner join user ON user.id = therad.user_id").
		Joins("left join ruang ON ruang.id = therad.ruang_id").
		Joins("left join likes ON likes.therad_id = therad.id").
		Where("status", "active").
		Where("therad.ruang_id", ruangId).
		Group("therad.id").
		Scan(&theradResponse)
	return repairTheradResponse(userId, theradResponse)
}

func GetSearchTheradByJudulAndIsi(key string, userId int) []model.TheradResponse {
	var theradResponse []model.TheradResponse
	database.DB.Model(&model.Therad{}).
		Select("therad.id", "therad.judul", "therad.user_id", "therad.isi", "therad.file", "therad.dilihat",
			"therad.status", "kategori_therad.kategori_name", "ruang.judul AS ruang_name",
			"CONCAT(user.firstname, ' ' ,user.lastname) AS author_name",
			"COUNT(likes.therad_id) AS total_like",
			"therad.created_at", "therad.updated_at").
		Joins("inner join kategori_therad ON kategori_therad.id = therad.kategori_therad_id").
		Joins("inner join user ON user.id = therad.user_id").
		Joins("left join ruang ON ruang.id = therad.ruang_id").
		Joins("left join likes ON likes.therad_id = therad.id").
		Where("status", "active").
		Where("therad.judul LIKE ? OR therad.isi LIKE ?", "%"+key+"%", "%"+key+"%").
		Group("therad.id").
		Scan(&theradResponse)
	return repairTheradResponse(userId, theradResponse)
}

func GetTheradByKategoriId(kategoriId, userId int) []model.TheradResponse {
	var theradResponse []model.TheradResponse
	database.DB.Model(&model.Therad{}).
		Select("therad.id", "therad.judul", "therad.user_id", "therad.isi", "therad.file", "therad.dilihat",
			"therad.status", "kategori_therad.kategori_name", "ruang.judul AS ruang_name",
			"CONCAT(user.firstname, ' ' ,user.lastname) AS author_name",
			"COUNT(likes.therad_id) AS total_like",
			"therad.created_at", "therad.updated_at").
		Joins("inner join kategori_therad ON kategori_therad.id = therad.kategori_therad_id").
		Joins("inner join user ON user.id = therad.user_id").
		Joins("left join ruang ON ruang.id = therad.ruang_id").
		Joins("left join likes ON likes.therad_id = therad.id").
		Where("status", "active").
		Where("therad.kategori_therad_id", kategoriId).
		Group("therad.id").
		Scan(&theradResponse)
	return repairTheradResponse(userId, theradResponse)
}

func GetTrendingTherad(userId int) []model.TheradResponse {
	var theradResponse []model.TheradResponse
	database.DB.Model(&model.Therad{}).
		Select("therad.id", "therad.judul", "therad.user_id", "therad.isi", "therad.file", "therad.dilihat",
			"therad.status", "kategori_therad.kategori_name", "ruang.judul AS ruang_name",
			"CONCAT(user.firstname, ' ' ,user.lastname) AS author_name",
			"COUNT(likes.therad_id) AS total_like",
			"therad.created_at", "therad.updated_at").
		Joins("inner join kategori_therad ON kategori_therad.id = therad.kategori_therad_id").
		Joins("inner join user ON user.id = therad.user_id").
		Joins("left join ruang ON ruang.id = therad.ruang_id").
		Joins("left join likes ON likes.therad_id = therad.id").
		Where("status", "active").
		Order("total_like desc").
		Group("therad.id").
		Scan(&theradResponse)
	return repairTheradResponse(userId, theradResponse)
}

func DeleteTherad(theradId int) {
	database.DB.Unscoped().Where("id", theradId).Delete(&model.Therad{})
}

func UpdateTherad(theradId int, therad model.Therad) {
	database.DB.Model(&model.Therad{}).Where("id", theradId).Updates(therad)
}
