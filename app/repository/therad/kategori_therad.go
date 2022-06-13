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
