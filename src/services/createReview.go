package services

import (
	"context"
	"time"

	"github.com/ArthurQR98/e-commerce/config"
	"github.com/ArthurQR98/e-commerce/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateReview(review models.Review) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := config.MongoCN.Database("ecommerce")
	col := db.Collection("reviews")

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
