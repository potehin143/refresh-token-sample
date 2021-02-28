package app

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type Parameters interface {
	ServerPort() string
	MongoUrl() string
	MongoUserName() string
	MongoPassword() string
	MongoDatabase() string
	TokenSecret() string
	AccessTokenExpiration() int64
	RefreshTokenExpiration() int64
}

type config struct {
	serverPort             string
	mongoHost              string
	mongoPort              string
	mongoUserName          string
	mongoPassword          string
	mongoDatabase          string
	tokenSecret            string
	accessTokenExpiration  int64 // nanoseconds
	refreshTokenExpiration int64 // nanoseconds
}

const (
	EMPTY                         = ""
	ServerPortDefault             = "8080"
	MongoHostDefault              = "localhost"
	MongoPortDefault              = "27017"
	MongoUsernameDefault          = "mongoroot"
	MongoPasswordDefault          = "mongopass"
	MongoDatabaseDefault          = "admin"
	TokenSecretDefault            = "0000-0000-0000-000"
	AccessTokenExpirationDefault  = 5 * 60 * 1e3       //milliseconds
	RefreshTokenExpirationDefault = 24 * 60 * 60 * 1e3 //milliseconds
)

func (conf config) ServerPort() string {
	return conf.serverPort
}

func (conf config) MongoUrl() string {
	return "mongodb://" + conf.mongoHost + ":" + conf.mongoPort
}

func (conf config) MongoUserName() string {
	return conf.mongoUserName
}

func (conf config) MongoPassword() string {
	return conf.mongoPassword
}

func (conf config) MongoDatabase() string {
	return conf.mongoDatabase
}

func (conf config) TokenSecret() string {
	return conf.tokenSecret
}

func (conf config) AccessTokenExpiration() int64 {
	return conf.accessTokenExpiration
}

func (conf config) RefreshTokenExpiration() int64 {
	return conf.refreshTokenExpiration
}

var instance *config
var once sync.Once

func GetParameters() Parameters {
	once.Do(func() {

		instance = &config{}

		serverPort := os.Getenv("SERVER_PORT")
		if serverPort != EMPTY {
			instance.serverPort = serverPort
		} else {
			log.Print("env variable SERVER_PORT is not set. Using default value ", ServerPortDefault)
			instance.serverPort = ServerPortDefault
		}

		mongoHost := os.Getenv("MONGO_HOST")
		if mongoHost != EMPTY {
			instance.mongoHost = mongoHost
		} else {
			log.Print("env variable MONGO_HOST is not set. Using default value ", MongoHostDefault)
			instance.mongoHost = MongoHostDefault
		}

		mongoPort := os.Getenv("MONGO_PORT")
		if mongoPort != EMPTY {
			instance.mongoPort = mongoPort
		} else {
			log.Print("env variable MONGO_PORT is not set. Using default value ", MongoPortDefault)
			instance.mongoPort = MongoPortDefault
		}

		mongoUserName := os.Getenv("MONGO_USERNAME")
		if mongoUserName != EMPTY {
			instance.mongoUserName = mongoUserName
		} else {
			log.Print("env variable MONGO_USERNAME is not set. Using default value")
			instance.mongoUserName = MongoUsernameDefault
		}

		mongoPassword := os.Getenv("MONGO_PASSWORD")
		if mongoPassword != EMPTY {
			instance.mongoPassword = mongoPassword
		} else {
			log.Print("env variable MONGO_PASSWORD is not set. Using default value")
			instance.mongoPassword = MongoPasswordDefault
		}

		mongoDatabase := os.Getenv("MONGO_DATABASE")
		if mongoDatabase != EMPTY {
			instance.mongoDatabase = mongoDatabase
		} else {
			log.Print("env variable MONGO_DATABASE is not set. Using default value ", MongoDatabaseDefault)
			instance.mongoDatabase = MongoDatabaseDefault
		}

		tokenSecret := os.Getenv("SECRET")
		if tokenSecret != EMPTY {
			instance.tokenSecret = tokenSecret
		} else {
			log.Print("env variable MONGO_PASSWORD is not set. Using default value")
			instance.tokenSecret = TokenSecretDefault
		}

		accessTokenExpirationMinutes, err := strconv.ParseInt(
			os.Getenv("ACCESS_TOKEN_EXPIRATION_MINUTES"),
			10,
			64)
		if err == nil {
			instance.accessTokenExpiration = accessTokenExpirationMinutes * 1e6
		} else {
			log.Print("env variable ACCESS_TOKEN_EXPIRATION_MINUTES is invalid or not set. Using default value ",
				AccessTokenExpirationDefault)
			instance.accessTokenExpiration = AccessTokenExpirationDefault * 1e6
		}

		refreshTokenExpiration, err := strconv.ParseInt(
			os.Getenv("REFRESH_TOKEN_EXPIRATION_MINUTES"),
			10,
			64)
		if err == nil {
			instance.refreshTokenExpiration = refreshTokenExpiration * 1e6
		} else {
			log.Print("env variable REFRESH_TOKEN_EXPIRATION_MINUTES is invalid or not set. Using default value ",
				RefreshTokenExpirationDefault)
			instance.refreshTokenExpiration = RefreshTokenExpirationDefault * 1e6
		}

	})
	return instance
}
