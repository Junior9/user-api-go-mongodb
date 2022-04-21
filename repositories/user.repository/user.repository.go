package userrepository

import (
	m "api/models"
	"context"
	"log"

	mongo "api/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = mongo.GetConnection("users")

var ctx = context.Background()

func Add(user m.User) (bool, error) {
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

func Get() ([]m.User, error) {
	filter := bson.D{}
	var users []m.User

	resp, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for resp.Next(ctx) {

		var user m.User

		err = resp.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetById(id string) (m.User, error) {

	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": oid}
	var user m.User

	userResp := collection.FindOne(ctx, filter)

	err := userResp.Decode(&user)

	if err != nil {
		return m.User{}, err
	}

	return user, nil
}
