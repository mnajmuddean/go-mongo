package main

import (
	"context"
	"go-mongodb/dbconnection"
	"go-mongodb/handlers"
	"go-mongodb/models"
)

func main() {

	client, serveErr := dbconnection.GetMongoClient("localhost", "27017")
	if serveErr != nil {
		panic(serveErr)
	}
	defer client.Disconnect(context.Background())

	// Call the CreatePerson function

	// err := handlers.CreatePerson(client, models.Person{})

	// if err != nil {
	// 	panic(err)
	// }

	// // Call the ReadPerson function

	// err := handlers.ReadPerson(client, models.Person{})
	// if err != nil {
	// 	panic(err)
	// }

	// // Call the UpdatePerson function
	// err := handlers.UpdatePerson(client, models.Person{})
	// if err != nil {
	// 	panic(err)
	// }

	// // Call the DeletePerson function
	// err := handlers.DeletePerson(client, models.Person{})
	// if err != nil {
	// 	panic(err)
	// }

	// // Call the UseQueryOperators function
	err := handlers.UseQueryOperators(client, models.Person{})
	if err != nil {
		panic(err)
	}

}
