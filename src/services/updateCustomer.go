package services

import (
	"context"
	"time"

	"github.com/ArthurQR98/e-commerce/config"
	"github.com/ArthurQR98/e-commerce/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateCustomer(customer models.Customer, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := config.MongoCN.Database("ecommerce")
	col := db.Collection("customers")

	register := make(map[string]interface{})
	if len(customer.Username) > 0 {
		register["username"] = customer.Username
	}
	if len(customer.Email) > 0 {
		register["email"] = customer.Email
	}
	if len(customer.First_name) > 0 {
		register["firstname"] = customer.First_name
	}
	if len(customer.Last_name) > 0 {
		register["lastname"] = customer.Last_name
	}
	if len(customer.Address) > 0 {
		register["address"] = customer.Address
	}
	if len(customer.Payments) > 0 {
		register["payment"] = customer.Payments
	}
	updateString := bson.M{"$set": register}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
