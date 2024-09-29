package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDatabase(username string, password string, address string, dbName string) *mongo.Database {
	var connStr string
	if len(username) > 0 && len(password) > 0 {
		connStr = fmt.Sprintf("mongodb://%s:%s@%s", username, password, address)
	} else {
		connStr = fmt.Sprintf("mongodb://%s", address)
	}

	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal("error while connecting to mongodb database: ", err)
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("failed to ping mongodb server: ", err)
	}

	log.Printf("successfully connected to Mongo database at %s", address)
	return mongoClient.Database(dbName)
}
