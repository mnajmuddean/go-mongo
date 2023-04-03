package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name    string `bson:"name,omitempty"`
	Age     int    `bson:"age,omitempty"`
	Address string `bson:"address, omitempty"`
}

func main() {

	////////////////////////////////////////// Connect to database ////////////////////////////////////////////////////////////

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // Create context that will expires in 10 seconds
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	/////////////////////////////////////////// End connection //////////////////////////////////////////////////////////

	// Get List Database :

	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(databases)

	/////////////////////////////////////////// Get collection instance //////////////////////////////////////////////

	collection := client.Database("myDB").Collection("Person")

	/////////////////////////////////////////// Insert Documents ////////////////////////////////////////////////

	// docs := []interface{}{
	// 	bson.D{{"name", "Ali"}, {"age", 23}, {"address", "Malaysia"}},
	// 	bson.D{{"name", "Abu"}, {"age", 24}, {"address", "Malaysia"}},
	// 	bson.D{{"name", "Khan"}, {"age", 25}, {"address", "Indonesia"}},
	// 	bson.D{{"name", "Lee"}, {"age", 23}, {"address", "Malaysia"}},
	// 	bson.D{{"name", "Kim"}, {"age", 24}, {"address", "Indonesia"}},
	// 	bson.D{{"name", "Siva"}, {"age", 23}, {"address", "India"}},
	// }

	// res, insertErr := collection.InsertMany(ctx, docs)
	// if insertErr != nil {
	// 	log.Fatal(insertErr)
	// }
	// fmt.Println(res)

	// cur, err := collection.Find(ctx, bson.D{})

	// if err != nil {
	// 	panic(err)
	// }
	// defer cur.Close(ctx)

	// var persons []Person
	// if err = cur.All(ctx, &persons); err != nil {
	// 	panic(err)
	// }

	// fmt.Println(persons)

	//////////////////////////////////////////// End Insert Documents //////////////////////////////////////////////////

	//////////////////////////////////////////// Retrieve a documents ///////////////////////////////////////////////

	// filter := bson.D{{"address", "Malaysia"}}

	// var result Person
	// err = collection.FindOne(context.TODO(), filter).Decode(&result)

	// if err != nil {
	// 	if err == mongo.ErrNoDocuments {
	// 		// This error means your query did not match any documents.
	// 		return
	// 	}
	// 	panic(err)
	// }

	// output, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", output)

	///////////////////////////////////////// End retrieve ///////////////////////////////////////////////////////////

	/////////////////////////////////////// Update Documents /////////////////////////////////////////////////////////

	// id, _ := primitive.ObjectIDFromHex("642a21516e5210b3daf9a7e2") //convert hexadecimal to objectID value
	// filter := bson.D{{"_id", id}}
	// update := bson.D{{"$set", bson.D{{"age", 30}}}}

	// result, err := collection.UpdateOne(context.TODO(), filter, update)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Documents updated: %v\n", result.ModifiedCount)
	/////////////////////////////////// End Update Documents ////////////////////////////////////////////////////////

	/////////////////////////////////// Delete  Documents ///////////////////////////////////////////////////////////

	// id, _ := primitive.ObjectIDFromHex("642a21516e5210b3daf9a7e6")
	// filter := bson.D{{"_id", id}}

	// result, err := collection.DeleteOne(context.TODO(), filter)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Documents deleted: %d\n", result.DeletedCount)

	/////////////////////////////////////// End Delete Documents ///////////////////////////////////////////////////////

	////////////////////////////////////// Using Query Operators /////////////////////////////////////////////////////////

	filter := bson.D{{"age", bson.D{{"$lte", 24}}}}

	result, err := collection.DeleteMany(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Documents deleted using query operators: %d\n", result.DeletedCount)

}
