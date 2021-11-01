package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Username   string             `bson:"username" json:"username,omitempty"`
	Email      string             `bson:"email" json:"email" valid:"email"`
	Password   string             `bson:"password" json:"password"`
	First_name string             `bson:"firstname" json:"firstname,omitempty"`
	Last_name  string             `bson:"lastname" json:"lastname,omitempty"`
	Address    []struct {
		Type    string `bson:"type" json:"type,omitempty"`
		Street  string `bson:"street" json:"street,omitempty"`
		City    string `bson:"city" json:"city,omitempty"`
		State   string `bson:"state" json:"state"`
		Pincode int64  `bson:"pincode" json:"pincode,omitempty"`
	}
	Payments string `bson:"payment" json:"payment,omitempty"`
}
