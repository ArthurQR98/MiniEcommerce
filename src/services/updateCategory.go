package services

import (
	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateCategory(category models.Category, ID string) (bool, error) {
	ctx, col, cancel := utils.ConnectDatabase("ecommerce", "category")
	defer cancel()

	register := make(map[string]interface{})
	if len(category.Name) > 0 {
		register["name"] = category.Name
	}
	if len(category.Description) > 0 {
		register["description"] = category.Description
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
