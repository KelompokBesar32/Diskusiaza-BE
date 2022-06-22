package profile

import (
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/database"
)

func GetDetailUser(id int) model.UserResponse {
	var userResponse model.UserResponse
	database.DB.Model(&model.User{}).Select("firstname", "lastname", "email", "nohp", "foto", "tanggal_lahir", "jenis_kelamin", "role.role_name").
		Joins("inner join role on role.id = user.role_id").
		Where("user.id", id).
		Scan(&userResponse)
	return userResponse
}

func UpdateUser(id int, mUser model.User) {
	database.DB.Model(&model.User{}).Where("id", id).Updates(mUser)
}
