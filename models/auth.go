package models

import (
	"../utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
)

/*
Структура прав доступа JWT
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

//структура для учётной записи пользователя
type LoginCredentials struct {
	Id       uint   `json:"id"`
	Password string `json:"password"`
}

//структура для учётной записи пользователя
type RefreshData struct {
	Id           uint   `json:"id"`
	RefreshToken string `json:"refreshToken"`
}

type Authentication struct {
	Id           uint   `json:"id"`
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}

func Login(id uint, password string) (map[string]interface{}, error) {

	authentication := &Authentication{}
	//err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	//if err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		return u.Message(false, "Email address not found")
	//	}
	//	return u.Message(false, "Connection error. Please retry")
	//}

	//err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	//if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Пароль не совпадает!!
	//	return u.Message(false, "Invalid login credentials. Please try again")
	//}
	//Работает! Войти в систему

	if id == 13 {
		return utils.Message(utils.BadCredentials), errors.New("bad credentials")
	}

	authentication.Id = id
	authentication.AccessToken = createAccessToken(id) // Сохраните токен в ответе
	authentication.RefreshToken = createRefreshToken(id)

	resp := utils.Message(utils.SUCCESS)
	resp["authentication"] = authentication
	return resp, nil
}
func Refresh(userId uint, refreshToken string) map[string]interface{} {
	authentication := &Authentication{}
	authentication.Id = userId
	authentication.AccessToken = createAccessToken(userId)
	authentication.RefreshToken = createRefreshToken(userId)

	if validate(refreshToken) {
		resp := utils.Message(utils.SUCCESS)
		resp["authentication"] = authentication
		return resp
	} else {
		resp := utils.Message("is invalid")
		return resp
	}

}

func createAccessToken(userId uint) string {
	//Создать токен JWT
	tk := &Token{UserId: userId}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET")))
	return tokenString
}

func createRefreshToken(userId uint) string {
	return "eoiuewroerwiutewropiuwerteoriu"
}

func validate(refreshToken string) bool {
	return true
}
