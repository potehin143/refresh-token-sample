package helper

import (
	"../app"
	"../repository"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

func createAccessToken(userId string, expiration int64) (string, error) {
	//creating JWT access token

	tokenClaims := jwt.MapClaims{}
	tokenClaims["userId"] = userId
	tokenClaims["expiresAt"] = expiration

	token := jwt.NewWithClaims(jwt.GetSigningMethod(SignMethod), tokenClaims)
	tokenString, err := token.SignedString([]byte(app.GetParameters().TokenSecret()))
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return tokenString, err
}

func createRefreshToken(userId string, expiration int64) (string, error) {
	//creating JWT refresh token
	tokenIdHex, err := repository.CreateRefreshToken(userId)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	tokenClaims := jwt.MapClaims{}
	tokenClaims["userId"] = userId
	tokenClaims["expiresAt"] = expiration
	tokenClaims["id"] = tokenIdHex

	token := jwt.NewWithClaims(jwt.GetSigningMethod(SignMethod), tokenClaims)
	tokenString, err := token.SignedString([]byte(app.GetParameters().TokenSecret()))

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return tokenString, err
}

func getAccessExpiration() int64 {
	return time.Now().Add(time.Duration(app.GetParameters().AccessTokenExpiration())).UnixNano()
}

func getRefreshExpiration() int64 {
	return time.Now().Add(time.Duration(app.GetParameters().RefreshTokenExpiration())).UnixNano()
}

func retrieveAccessToken(tokenString string) (string, error) {
	tokenClaims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.GetParameters().TokenSecret()), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("token is invalid")
	}
	expiresAt, ok := tokenClaims["expiresAt"].(float64) // It was parsed to tokenClaims as float64
	if !ok {
		return "", errors.New("can't parse expiresAt")
	}
	if int64(expiresAt) < time.Now().UnixNano() { //Token is expired
		return "", errors.New("token is expired")
	}

	tokenUserId, ok := tokenClaims["userId"].(string)
	if !ok {
		return "", errors.New("can't parse userId")
	}

	return tokenUserId, nil
}

func retrieveRefreshToken(tokenString string) (string, string, error) {
	tokenClaims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.GetParameters().TokenSecret()), nil
	})

	if err != nil {
		return "", "", err
	}

	if !token.Valid {
		return "", "", errors.New("token is invalid")
	}
	expiresAt, ok := tokenClaims["expiresAt"].(float64) // It was parsed to tokenClaims as float64
	if !ok {
		return "", "", errors.New("can't parse expiresAt")
	}
	if int64(expiresAt) < time.Now().UnixNano() { //Token is expired
		return "", "", errors.New("token is expired")
	}

	tokenUserId, ok := tokenClaims["userId"].(string)
	if !ok {
		return "", "", errors.New("can't parse userId")
	}
	tokenIdHex, ok := tokenClaims["id"].(string)
	if !ok {
		return tokenUserId, "", errors.New("cant parse userId")
	}
	return tokenUserId, tokenIdHex, nil
}
