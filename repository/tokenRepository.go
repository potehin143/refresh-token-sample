package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

const TOKENS = "tokens"

func CreateRefreshToken(clientId string) (string, error) {

	collection, err := GetCollection(TOKENS)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	token := RefreshToken{UserId: clientId, ExpirationDate: time.Now()}

	insertResult, err2 := collection.InsertOne(context.TODO(), token)
	if err2 != nil {
		log.Fatal(err2)
		return "", err2
	}

	oid, _ := insertResult.InsertedID.(primitive.ObjectID)

	return oid.Hex(), nil
}

func DeleteToken(hex string, userId string) (int64, error) {
	collection, err := GetCollection(TOKENS)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	objID, _ := primitive.ObjectIDFromHex(hex)

	deleteResult, _ := collection.DeleteOne(context.TODO(), bson.M{"_id": objID, "userid": userId})
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	if deleteResult == nil {
		return 0, err
	}

	return deleteResult.DeletedCount, err
}

func DeleteTokens(userId string) (int64, error) {
	collection, err := GetCollection(TOKENS)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	deleteResult, _ := collection.DeleteMany(context.TODO(), bson.M{"userid": userId})
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	if deleteResult == nil {
		return 0, err
	}

	return deleteResult.DeletedCount, err
}

type RefreshToken struct {
	UserId         string
	ExpirationDate time.Time //TODO remove it
}
