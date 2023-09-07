package storage

import (
	"context"
	"fmt"
	"log"

	osutils "github.com/krissolui/go-utils/os-utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultMongoURL       = "mongodb://mongo:27017"
	defualtDatabaseName   = "sessions"
	defualtCollectionName = "sessions"
)

func NewStorage() (Storage, error) {
	client, err := connectMongo()
	if err != nil {
		return Storage{}, err
	}
	// defer client.Disconnect(context.TODO())

	collection, err := getCollection(client)
	if err != nil {
		return Storage{}, err
	}

	storage := Storage{
		Client:     client,
		Collection: collection,
	}
	return storage, nil
}

func connectMongo() (*mongo.Client, error) {
	url := osutils.GetEnv("MONGO_URL", defaultMongoURL)

	clientOptions := options.Client()
	clientOptions.ApplyURI(url)
	clientOptions.SetAuth(options.Credential{
		Username: osutils.GetEnv("MONGO_USERNAME"),
		Password: osutils.GetEnv("MONGO_PASSWORD"),
	})

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Printf("Failed to connect MongoDB at %s: %v\n", url, err)
		return nil, err
	}

	log.Printf("Connected to MongoDB at %s\n", url)
	return client, nil
}

func getCollection(client *mongo.Client) (*mongo.Collection, error) {
	databaseName := osutils.GetEnv("MONGO_DATABASE", defualtDatabaseName)
	db := client.Database(databaseName)
	if db == nil {
		return nil, fmt.Errorf("failed to get database %s", databaseName)
	}

	collectionName := osutils.GetEnv("MONGO_COLLECTION", defualtCollectionName)
	collection := db.Collection(collectionName)
	if collection == nil {
		return nil, fmt.Errorf("failed to get collection %s", collectionName)
	}

	log.Printf("Retrieved handler for [database, collection]: [%s, %s]", databaseName, collectionName)

	index := mongo.IndexModel{
		Keys:    bson.D{{Key: "expire_at", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(0),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), index)
	if err != nil {
		return nil, err
	}

	return collection, nil
}
