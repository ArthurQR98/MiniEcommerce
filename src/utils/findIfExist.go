package utils

import (
	"os"

	"github.com/ArthurQR98/e-commerce/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var DBname = os.Getenv("MONGO_DB")

func FindIfExistCategory(id primitive.ObjectID) (bool, string) {
	ctx, col, cancel := ConnectDatabase(DBname, "category")
	defer cancel()

	condition := bson.M{"_id": id}
	var category models.Category

	err := col.FindOne(ctx, condition).Decode(&category)
	if err != nil {
		return false, "Not found"
	}
	return true, ""
}

func FindIfExistReview(id primitive.ObjectID) (bool, string) {
	ctx, col, cancel := ConnectDatabase(DBname, "reviews")
	defer cancel()

	condition := bson.M{"_id": id}
	var review models.Review

	err := col.FindOne(ctx, condition).Decode(&review)
	if err != nil {
		return false, "Not found"
	}
	return true, ""
}

func FindIfExistProduct(id primitive.ObjectID) (bool, string) {
	ctx, col, cancel := ConnectDatabase(DBname, "products")
	defer cancel()

	condition := bson.M{"_id": id}
	var product models.Product

	err := col.FindOne(ctx, condition).Decode(&product)
	if err != nil {
		return false, "Not found"
	}
	return true, ""
}
