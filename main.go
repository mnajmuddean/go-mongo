package main

import (
	"go-mongodb/handlers"
	"go-mongodb/models"
)

func main() {

	err := handlers.CreatePerson(models.Person{})

	if err != nil {
		panic(err)
	}

	// err := handlers.ReadPerson(models.Person{})
	// if err != nil {
	// 	panic(err)
	// }

	// // Call the updatePerson function
	// err := handlers.UpdatePerson(models.Person{})
	// if err != nil {
	// 	panic(err)
	// }

	// // Call the deletePerson function
	// err := handlers.DeletePerson(models.Person{})
	// if err != nil {
	// 	panic(err)
	// }

	// // Call the useQueryOperators function
	// err := handlers.UseQueryOperators(models.Person{})
	// if err != nil {
	// 	panic(err)
	// }

}
