package services

import (
	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateOrder(order models.Order, ID string) (bool, error) {
	ctx, col, cancel := utils.ConnectDatabase(DBname, "orders")
	defer cancel()
	order.Items.Quantity = len(order.Items.Product)
	updateString := bson.M{"$set": order}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
