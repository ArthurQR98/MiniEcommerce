package services

import (
	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrder(order models.Order) (string, bool, error) {
	ctx, col, cancel := utils.ConnectDatabase("ecommerce", "orders")
	defer cancel()

	result, err := col.InsertOne(ctx, order)
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
