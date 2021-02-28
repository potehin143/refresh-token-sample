package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const USERS = "users"

func GetUser(userId string) (*User, bool, error) {
	var result *User

	collection, err := GetCollection(USERS)
	if err != nil {
		log.Fatal(err)
		return result, false, err
	}

	filter := bson.D{{"userid", userId}}
	findOptions := options.Find()
	findOptions.SetLimit(1)

	cur, err2 := collection.Find(context.TODO(), filter, findOptions)
	if err2 != nil {
		log.Fatal(err2)
		return result, false, err2
	}
	exists := false
	for cur.Next(context.TODO()) {
		var elem User
		err3 := cur.Decode(&elem)
		if err3 != nil {
			log.Fatal(err3)
			return result, exists, err
		}
		result = &elem
		exists = true
	}
	_ = cur.Close(context.TODO())

	return result, exists, nil
}

func CreateUser(userId string, passwordHash string) (string, error) {

	collection, err := GetCollection(USERS)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	user := User{UserId: userId, PasswordHash: passwordHash}

	insertResult, err2 := collection.InsertOne(context.TODO(), user)
	if err2 != nil {
		log.Fatal(err2)
		return "", err2
	}

	oid, _ := insertResult.InsertedID.(primitive.ObjectID)

	return oid.Hex(), nil
}

type User struct {
	UserId       string
	PasswordHash string
}
