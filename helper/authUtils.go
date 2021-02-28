package helper

import (
	"../repository"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	//	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"log"
	"net/http"
)

const (
	SignMethod = "HS512"
)

/*
Структура прав доступа JWT
*/
type TokenClaims struct {
	UserId string
	jwt.StandardClaims
}

//структура для учётной записи пользователя
type LoginCredentials struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

//структура для учётной записи пользователя
type RefreshData struct {
	Id           string `json:"id"`
	RefreshToken string `json:"refreshToken"`
}

type Authentication struct {
	Id           string `json:"id"`
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}

func Register(userId string, password string, writer http.ResponseWriter) (map[string]interface{}, error) {

	_, exists, _ := repository.GetUser(userId)

	if exists {
		response := Message(AlreadyExists)
		writer.WriteHeader(http.StatusConflict)
		return response, errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		response := Message(InternalServerError)
		writer.WriteHeader(http.StatusInternalServerError)
		return response, err
	}

	_, _ = repository.CreateUser(userId, string(hash))
	resp := Message(SUCCESS)
	return resp, nil
}

func Login(userId string, password string, writer http.ResponseWriter) (map[string]interface{}, error) {

	user, exists, err := repository.GetUser(userId)

	if err != nil {
		response := Message(InternalServerError)
		writer.WriteHeader(http.StatusInternalServerError)
		return response, err
	}

	if !exists {
		response := Message(BadCredentials)
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return response, errors.New("bad credentials")
	}

	passwordHash := user.PasswordHash

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		response := Message(BadCredentials)
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return response, errors.New("bad credentials")
	}

	_, err2 := repository.DeleteTokens(userId)
	if err2 != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		response := Message(InternalServerError)
		return response, err2
	}

	authentication, err := createAuthentication(userId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		response := Message(InternalServerError)
		return response, err
	}

	resp := Message(SUCCESS)
	resp["authentication"] = authentication
	return resp, nil
}

func Logout(writer http.ResponseWriter, request *http.Request) (map[string]interface{}, error) {
	userId := request.Context().Value("userId").(string)
	_, err := repository.DeleteTokens(userId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		response := Message(InternalServerError)
		return response, err
	}

	resp := Message(SUCCESS)
	return resp, nil
}

func Refresh(userId string, refreshToken string, writer http.ResponseWriter) (map[string]interface{}, error) {

	/*
		tokenClaims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(refreshToken, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(app.GetParameters().TokenSecret()), nil
		})

		if err != nil{
			writer.WriteHeader(http.StatusUnauthorized)
			resp := Message(Unauthorized)
			return resp, err
		}

		if !token.Valid{
			resp := Message(Unauthorized)
			writer.WriteHeader(http.StatusUnauthorized)
			return resp, errors.New("token is invalid")
		}
		expiresAt, ok := tokenClaims["expiresAt"].(float64) // It was parsed to tokenClaims as float64
		if !ok {
			resp := Message(Unauthorized)
			writer.WriteHeader(http.StatusUnauthorized)
			return resp, errors.New("can't parse expiresAt")
		}
		if int64(expiresAt) < time.Now().UnixNano(){ //Token is expired
			resp := Message(Unauthorized)
			writer.WriteHeader(http.StatusUnauthorized)
			return resp, errors.New("token is expired")
		}

		tokenUserId, ok := tokenClaims["userId"].(string)
		if !ok {
			resp := Message(Unauthorized)
			writer.WriteHeader(http.StatusUnauthorized)
			return resp, errors.New("can't parse userId")
		}
		tokenIdHex, ok := tokenClaims["id"].(string)
		if !ok {
			log.Println("cant parse userId")
			resp := Message(Unauthorized)
			writer.WriteHeader(http.StatusUnauthorized)
			return resp, nil
		}

	*/
	tokenUserId, tokenIdHex, err := retrieveRefreshToken(refreshToken)

	if err != nil {
		resp := Message(Unauthorized)
		writer.WriteHeader(http.StatusUnauthorized)
		return resp, err
	}

	if userId != tokenUserId {
		resp := Message(Forbidden)
		writer.WriteHeader(http.StatusForbidden)
		return resp, err
	}

	count, err2 := repository.DeleteToken(tokenIdHex, tokenUserId)

	if err2 != nil {
		resp := Message(InternalServerError)
		writer.WriteHeader(http.StatusInternalServerError)
		return resp, nil
	}

	if count == 0 { //No any tokens were found to delete, so it already used or not belongs to this user
		writer.WriteHeader(http.StatusUnauthorized)
		resp := Message(Unauthorized)
		return resp, nil
	}

	authentication, err := createAuthentication(userId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		response := Message(InternalServerError)
		return response, err
	}

	resp := Message(SUCCESS)
	resp["authentication"] = authentication
	return resp, nil

}

func createAuthentication(userId string) (*Authentication, error) {

	accessToken, err := createAccessToken(userId, getAccessExpiration())
	if err != nil {
		log.Print(err)
		return nil, err
	}

	refreshToken, err := createRefreshToken(userId, getRefreshExpiration())
	if err != nil {
		log.Print(err)
		return nil, err
	}

	authentication := &Authentication{}
	authentication.Id = userId
	authentication.AccessToken = accessToken
	authentication.RefreshToken = refreshToken
	return authentication, nil
}
