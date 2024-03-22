package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Collections *mongo.Collection

func ConnectDB() {
	mongoConnStr := os.Getenv("MONGO_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoConnStr))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	Collections = client.Database("sitemap").Collection("companies")
	count, err := Collections.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		fmt.Println(err, count)
		panic(err)
	}
	fmt.Println("Mongo connection established!")
}
