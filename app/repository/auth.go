package repository

import (
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/database"
	"golang.org/x/crypto/bcrypt"
)

func CreateUsers(user model.User) {
	// set role id 1, because users id 1 is normal user
	user.RoleID = 1
	// hash password and set password hash
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(passwordHash)
	database.DB.Create(&user)
}

func GetUsersByEmail(email string) bool {
	mUser := model.User{}
	database.DB.First(&mUser, "email = ?", email)
	if mUser.Email == "" {
		return true
	}
	return false
}

func Login(user model.User) (model.User, bool) {
	mUsers := model.User{}
	emailErr := database.DB.Where("email = ?", user.Email).First(&mUsers).Error
	if emailErr != nil {
		return mUsers, false
	}
	passErr := bcrypt.CompareHashAndPassword([]byte(mUsers.Password), []byte(user.Password))
	if passErr != nil && passErr == bcrypt.ErrMismatchedHashAndPassword {
		return mUsers, false
	}
	return mUsers, true
}

func UpdateToken(token string, userId uint) {
	database.DB.Model(&model.User{}).Where("id = ?", userId).Update("token", token)
}

func CheckTokenFromDB(token string) bool {
	mUser := model.User{}
	database.DB.First(&mUser, "token = ?", token)
	if mUser.Token == "" {
		return false
	}
	return true
}

func RemoveToken(token string) {
	database.DB.Model(&model.User{}).Where("token = ?", token).Update("token", "")
}
