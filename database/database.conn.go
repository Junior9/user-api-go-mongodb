package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection(collection string) *mongo.Collection {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27018"))

	if err != nil {
		fmt.Println("Error 1")
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error 3")
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx)
	return client.Database("api_go").Collection(collection)
}
