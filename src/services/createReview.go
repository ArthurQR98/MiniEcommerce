package services

import (
	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateReview(review models.Review) (string, bool, error) {
	ctx, col, cancel := utils.ConnectDatabase("ecommerce", "reviews")
	defer cancel()

	register := bson.M{
		"postDate":   review.PostDate,
		"title":      review.Title,
		"body":       review.Body,
		"customerId": review.Customer,
		"rating":     review.Rating,
	}
	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
