package services

import (
	"github.com/ArthurQR98/e-commerce/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteProduct(ID string) error {
	ctx, col, cancel := utils.ConnectDatabase(DBname, "products")
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{"_id": objID}
	_, err := col.DeleteOne(ctx, condition)
	return err
}
