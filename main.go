package main

import (
	"context"
	"go-mongodb/dbconnection"
	"go-mongodb/handlers"
	"go-mongodb/models"
)

func main() {

	// // Get only one client (GetMongoClient)

	client, err := dbconnection.GetMongoClient("localhost", "27017")
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	// Call the CreatePerson function

	handlers.CreatePerson(client, models.Person{})

	// Call the ReadPerson function

	handlers.ReadPerson(client, models.Person{})

	// // Call the UpdatePerson function
	handlers.UpdatePerson(client, models.Person{})

	// // Call the DeletePerson function
	handlers.DeletePerson(client, models.Person{})

	// Call the UseQueryOperators function
	handlers.UseQueryOperators(client, models.Person{})

}
