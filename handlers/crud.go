package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go-mongodb/dbconnection"
	"go-mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Person struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Age     int                `json:"age,omitempty" bson:"age,omitempty"`
	Country string             `json:"country,omitempty" bson:"country,omitempty"`
}

func CreatePerson(models.Person) error {
	// Get a connection to the database
	client, err := dbconnection.GetMongoClient()
	if err != nil {
		panic(err)
	}

	// Get a handle to the "people" collection in the "test" database
	collection := client.Database("myDB").Collection("Person")

	docs := []interface{}{
		Person{Name: "Muhammad Najmuddin", Age: 23, Country: "Malaysia"},
		Person{Name: "Bryan", Age: 24, Country: "Malaysia"},
		Person{Name: "Ahmas", Age: 25, Country: "Indonesia"},
		Person{Name: "Shah", Age: 26, Country: "Singapore"},
		Person{Name: "Wong", Age: 23, Country: "Thailand"},
		Person{Name: "Ali", Age: 24, Country: "Malaysia"},
		Person{Name: "Raj", Age: 25, Country: "India"},
		Person{Name: "Tan", Age: 26, Country: "China"},
		Person{Name: "Kim", Age: 24, Country: "Hong Kong"},
		Person{Name: "Lee", Age: 25, Country: "Hong Kong"},
	}

	ctx := context.Background()

	res, insertErr := collection.InsertMany(ctx, docs)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Println(res)

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		panic(err)
	}
	defer cur.Close(ctx)

	var persons []Person
	if err = cur.All(ctx, &persons); err != nil {
		panic(err)
	}

	fmt.Println(persons)

	return nil
}

func ReadPerson(models.Person) error {
	// Get a connection to the database
	client, err := dbconnection.GetMongoClient()
	if err != nil {
		panic(err)
	}

	// Get a handle to the "people" collection in the "test" database
	collection := client.Database("myDB").Collection("Person")

	filter := bson.D{{"country", "Malaysia"}}

	var result Person
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return err
		}
		panic(err)
	}

	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)

	return nil
}

func UpdatePerson(models.Person) error {
	// Get a connection to the database
	client, err := dbconnection.GetMongoClient()
	if err != nil {
		panic(err)
	}

	// Get a handle to the "people" collection in the "test" database
	collection := client.Database("myDB").Collection("Person")

	id, _ := primitive.ObjectIDFromHex("642a888c74c81a173a34883a") //convert hexadecimal to objectID value
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"age", 30}}}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)

	return nil
}

func DeletePerson(models.Person) error {
	// Get a connection to the database
	client, err := dbconnection.GetMongoClient()
	if err != nil {
		panic(err)
	}

	// Get a handle to the "people" collection in the "test" database
	collection := client.Database("myDB").Collection("Person")
	id, _ := primitive.ObjectIDFromHex("642a888c74c81a173a34883a")
	filter := bson.D{{"_id", id}}

	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Documents deleted: %d\n", result.DeletedCount)

	return nil
}

func UseQueryOperators(models.Person) error {
	// Get a connection to the database
	client, err := dbconnection.GetMongoClient()
	if err != nil {
		panic(err)
	}

	// Get a handle to the "people" collection in the "test" database
	collection := client.Database("myDB").Collection("Person")

	filter := bson.D{{"age", bson.D{{"$lte", 24}}}}

	result, err := collection.DeleteMany(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Documents deleted using query operators: %d\n", result.DeletedCount)

	return nil
}
