package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

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

func CreatePerson(client *mongo.Client, person models.Person) error {

	// Get a handle to the "Person" collection in the "myDB" database
	collection := client.Database("myDB").Collection("Person")

	docs := []interface{}{
		Person{Name: "Muhammad Najmuddin", Age: 23, Country: "Malaysia"},
		Person{Name: "Bryan", Age: 24, Country: "Malaysia"},
		Person{Name: "Ahmad", Age: 25, Country: "Indonesia"},
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

	jsonBytes, err := json.MarshalIndent(persons, "", "    ")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonBytes)

	return nil
}

func ReadPerson(client *mongo.Client, person models.Person) error {

	// Get a handle to the "Person" collection in the "myDB" database
	collection := client.Database("myDB").Collection("Person")

	filter := bson.D{{Key: "country", Value: "Hong Kong"}}

	var result Person
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return err
		}
		panic(err)
	}

	fmt.Println("Below is the person details that is from Hong Kong : ")

	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)

	return nil
}

func UpdatePerson(client *mongo.Client, person models.Person) error {

	// Get a handle to the "Person" collection in the "myDB" database
	collection := client.Database("myDB").Collection("Person")

	//convert hexadecimal to objectID value
	filter := bson.D{{Key: "name", Value: "Ali"}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "age", Value: 30}}}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}

	// Find the updated document
	var updatedDoc Person
	err = collection.FindOne(context.TODO(), filter).Decode(&updatedDoc)
	if err != nil {
		panic(err)
	}

	fmt.Println("Below is the document that have been updated :")
	// Print the updated document in JSON format
	output, err := json.MarshalIndent(updatedDoc, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)

	return nil
}

func DeletePerson(client *mongo.Client, person models.Person) (interface{}, error) {

	// Get a handle to the "Person" collection in the "myDB" database
	collection := client.Database("myDB").Collection("Person")
	filter := bson.D{{Key: "name", Value: "Muhammad Najmuddin"}}

	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		fmt.Println("Delete Status : Failed")
		panic(err)
	}

	fmt.Println("Delete Status : Success")
	fmt.Printf("Documents successfully deleted: %d\n", result.DeletedCount)

	return nil, nil
}

func UseQueryOperators(client *mongo.Client, person models.Person) error {

	// Get a handle to the "Person" collection in the "myDB" database
	collection := client.Database("myDB").Collection("Person")

	filter := bson.D{{Key: "age", Value: bson.D{{Key: "$lte", Value: 25}}}}

	result, err := collection.CountDocuments(context.TODO(), filter)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Documents count using query operators: %d\n", result)

	return nil
}
