package database

import (
	"context"
	"extinct-leaderboard/models"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

const databaseName = "extinct"
const collectionName = "users"

func Init(uri string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetClient() *mongo.Client {
	return client
}

func GetDatabase() *mongo.Database {
	return client.Database(databaseName)
}

func GetLeaderboard(key string) []models.User {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{key, -1}})

	if client == nil {
		panic("client has not been initialized")
	}

	coll := GetDatabase().Collection(collectionName)
	cursor, err := coll.Find(context.TODO(), filter, opts)

	if err != nil {
		panic(err)
	}

	var result []models.User

	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found")
		return nil
	}

	if err != nil {
		panic(err)
	}

	return result
}
