package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Age     int                `json:"age,omitempty" bson:"age,omitempty"`
	Country string             `json:"country,omitempty" bson:"country,omitempty"`
}
