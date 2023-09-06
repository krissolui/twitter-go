package storage

import (
	"context"
	"fmt"
	"log"
	"session-service/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultMongoURL       = "mongodb://mongo:27017"
	defualtDatabaseName   = "sessions"
	defualtCollectionName = "sessions"
)

func ConnectMongo() (*mongo.Client, error) {
	url := utils.GetEnvOrDefault("MONGO_URL", defaultMongoURL)

	clientOptions := options.Client()
	clientOptions.ApplyURI(url)
	clientOptions.SetAuth(options.Credential{
		Username: utils.GetEnvOrDefault("MONGO_USERNAME"),
		Password: utils.GetEnvOrDefault("MONGO_PASSWORD"),
	})

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Printf("Failed to connect MongoDB at %s: %v\n", url, err)
		return nil, err
	}

	log.Printf("Connected to MongoDB at %s\n", url)
	return client, nil
}

func GetCollection(client *mongo.Client) (*mongo.Collection, error) {
	databaseName := utils.GetEnvOrDefault("MONGO_DATABASE", defualtDatabaseName)
	db := client.Database(databaseName)
	if db == nil {
		return nil, fmt.Errorf("failed to get database %s", databaseName)
	}

	collectionName := utils.GetEnvOrDefault("MONGO_COLLECTION", defualtCollectionName)
	collection := db.Collection(collectionName)
	if collection == nil {
		return nil, fmt.Errorf("failed to get collection %s", collectionName)
	}

	log.Printf("Retrieved handler for [database, collection]: [%s, %s]", databaseName, collectionName)
	return collection, nil
}
