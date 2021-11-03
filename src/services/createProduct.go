package services

import (
	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(product models.Product) (string, bool, error) {
	ctx, col, cancel := utils.ConnectDatabase("ecommerce", "products")
	defer cancel()

	result, err := col.InsertOne(ctx, product)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
