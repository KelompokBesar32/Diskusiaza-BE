package middleware

import (
	"Diskusiaza-BE/app/model"
	"Diskusiaza-BE/constants"
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateTokenUsers(user model.User) (string, error) {
	claims := &jwt.MapClaims{
		"data": map[string]interface{}{
			"firstname": user.Firstname,
			"lastname":  user.Lastname,
			"email":     user.Email,
			"id":        user.ID,
			"roleId":    user.RoleID,
		},
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.ScreetJwtForUser))
}

func DecodeTokenUsers(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	return claims, err
}
