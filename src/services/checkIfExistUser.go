package services

import (
	"os"

	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/utils"
	"go.mongodb.org/mongo-driver/bson"
)

var DBname = os.Getenv("MONGO_DB")

func CheckIfExistUser(email string) (models.Customer, bool, string) {
	ctx, col, cancel := utils.ConnectDatabase(DBname, "customers")
	defer cancel()
	condition := bson.M{"email": email}
	var result models.Customer

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
