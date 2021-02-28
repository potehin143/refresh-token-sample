package repository

import (
	"../app"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func GetClient() (*mongo.Client, error) {
	// Setting client options
	credentials := options.Credential{
		Username:      app.GetParameters().MongoUserName(),
		Password:      app.GetParameters().MongoPassword(),
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    app.GetParameters().MongoDatabase(),
	}

	clientOptions := options.Client().ApplyURI(app.GetParameters().MongoUrl()).SetAuth(credentials)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Print(err)
	}
	return client, err
}

func GetCollection(collectionName string) (*mongo.Collection, error) {
	client, err := GetClient()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client.Database(app.GetParameters().MongoDatabase()).Collection(collectionName), err
}
