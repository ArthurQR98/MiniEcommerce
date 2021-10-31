package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	User  string             `bson:"userId" json:"userId"`
	State string             `bson:"state" json:"state"`
	Item  struct {
		Product  []Product `bson:"product" json:"product"`
		Quantity int       `bson:"quantity" json:"quantity"`
		Cost     int       `bson:"cost" json:"cost"`
	}
}
