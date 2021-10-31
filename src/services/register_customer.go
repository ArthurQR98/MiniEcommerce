package services

import (
	"context"
	"time"

	"github.com/ArthurQR98/e-commerce/config"
	"github.com/ArthurQR98/e-commerce/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterCustomer(customer models.Customer) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := config.MongoCN.Database("ecommerce")
	col := db.Collection("customers")

	customer.Password, _ = EncryptPassword(customer.Password)
	result, err := col.InsertOne(ctx, customer)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
