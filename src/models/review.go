package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	PostDate time.Time          `bson:"postDate" json:"postDate,omitempty"`
	Title    string             `bson:"title" json:"title"`
	Body     string             `bson:"body" json:"body"`
	Customer primitive.ObjectID `bson:"customerId,omitempty" json:"customerId,omitempty"`
	Rating   int                `bson:"rating" json:"rating,omitempty"`
}
