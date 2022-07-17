package therad

import (
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/database"
)

func GetKategoriTherad() []model.KategoriTheradResponse {
	var res []model.KategoriTheradResponse
	database.DB.Model(&model.KategoriTherad{}).Scan(&res)
	return res
}

func GetKategoriTheradById(kategoriId int) model.KategoriTheradResponse {
	var res model.KategoriTheradResponse
	database.DB.First(&model.KategoriTherad{}).Where("id", kategoriId).
		Scan(&res)
	return res
}
