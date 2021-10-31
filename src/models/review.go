package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Product  string             `bson:"productId" json:"productId"`
	PostDate time.Time          `bson:"postDate" json:"postDate"`
	Title    string             `bson:"title" json:"title"`
	Body     string             `bson:"body" json:"body"`
	Customer Customer           `bson:"customer" json:"customer"`
	Rating   int                `bson:"rating" json:"rating,omitempty"`
}
