package services

import (
	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateReview(review models.Review, ID string) (bool, error) {
	ctx, col, cancel := utils.ConnectDatabase("ecommerce", "reviews")
	defer cancel()
	register := make(map[string]interface{})
	if len(review.Title) > 0 {
		register["title"] = review.Title
	}
	if len(review.Body) > 0 {
		register["body"] = review.Body
	}
	register["rating"] = review.Rating
	updateString := bson.M{"$set": register}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
