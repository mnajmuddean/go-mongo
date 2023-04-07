package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestInsert(t *testing.T) {
	client := createMockClient()
	col := client.Database("test").Collection("City")

	doc := bson.D{
		{Key: "city_name", Value: "Muar"},
		{Key: "state", Value: "Johor"},
	}

	_, err := col.InsertOne(context.Background(), doc)
	assert.NoError(t, err)

	result := col.FindOne(context.Background(), bson.D{{Key: "city_name", Value: "Muar"}})
	assert.NoError(t, result.Err())

	var retrievedDoc bson.M
	err = result.Decode(&retrievedDoc)
	assert.NoError(t, err)
	assert.Equal(t, "Muar", retrievedDoc["city_name"])
	assert.Equal(t, "Johor", retrievedDoc["state"])
}

func createMockClient() *mongo.Client {
	uri := "mongodb://localhost:27017"
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(opts)
	if err != nil {
		panic(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	return client

}
