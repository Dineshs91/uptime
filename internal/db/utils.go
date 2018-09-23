package db

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// GenerateObjectID function generates & adds it to passed struct.
func GenerateObjectID() objectid.ObjectID {
	return objectid.New()
}

// CreateUser persists the user to db.
func CreateUser(user User) interface{} {
	dbClient := GetDbClient()
	collection := dbClient.Database("uptime").Collection("users")

	result, _ := collection.InsertOne(
		context.Background(),
		user,
	)
	return result.InsertedID
}

// GetUser from db.
func GetUser(email string) User {
	dbClient := GetDbClient()
	collection := dbClient.Database("uptime").Collection("users")

	user := User{}
	collection.FindOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("email", email),
		),
	).Decode(&user)

	return user
}

// AddMonitoringURL function persists the value in db.
func AddMonitoringURL(monitorURL MonitorURL) interface{} {
	dbClient := GetDbClient()
	collection := dbClient.Database("uptime").Collection("monitorURL")

	result, _ := collection.InsertOne(
		context.Background(),
		monitorURL,
	)
	return result.InsertedID
}

// GetMonitoringURL function gets monitor url from db.
func GetMonitoringURL() MonitorURL {
	dbClient := GetDbClient()
	collection := dbClient.Database("uptime").Collection("monitorURL")

	monitorURL := MonitorURL{}
	collection.FindOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("protocol", "https"),
		),
	).Decode(&monitorURL)

	return monitorURL
}