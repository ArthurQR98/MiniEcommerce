package services

import (
	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCategory(category models.Category) (string, bool, error) {
	ctx, col, cancel := utils.ConnectDatabase("ecommerce", "category")
	defer cancel()

	register := bson.M{
		"name":        category.Name,
		"description": category.Description,
	}
	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
