package services

import (
	"context"
	"time"

	"github.com/ArthurQR98/e-commerce/config"
	"github.com/ArthurQR98/e-commerce/src/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckIfExistUser(email string) (models.Customer, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := config.MongoCN.Database("ecommerce")
	col := db.Collection("customers")

	condition := bson.M{"email": email}
	var result models.Customer

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
