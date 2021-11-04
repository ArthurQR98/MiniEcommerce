package services

import (
	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateOrder(order models.Order, ID string) (bool, error) {
	ctx, col, cancel := utils.ConnectDatabase("ecommerce", "orders")
	defer cancel()

	register := make(map[string]interface{})
	if len(order.State) > 0 {
		register["state"] = order.State
	}
	if len(order.Items.Product) > 0 {
		register["items"] = order.Items
	}
	register["date"] = order.Date
	updateString := bson.M{"$set": register}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
