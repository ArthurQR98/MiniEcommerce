package services

import (
	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterCustomer(customer models.Customer) (string, bool, error) {
	ctx, col, cancel := utils.ConnectDatabase("ecommerce", "customers")
	defer cancel()

	customer.Password, _ = EncryptPassword(customer.Password)
	result, err := col.InsertOne(ctx, customer)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
