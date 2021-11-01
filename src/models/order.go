package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	User  primitive.ObjectID `bson:"userId" json:"userId"`
	State string             `bson:"state" json:"state"`
	Items struct {
		Product  []primitive.ObjectID `bson:"products" json:"products"`
		Quantity int                  `bson:"quantity" json:"quantity"`
	}
	Date time.Time `bson:"date" json:"date,omitempty"`
}
