package services

import (
	"context"
	"log"

	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProducts(paginate int64) ([]*models.Product, bool) {
	ctx, col, cancel := utils.ConnectDatabase(DBname, "products")
	defer cancel()

	var results []*models.Product
	option := options.Find()
	option.SetLimit(20)
	option.SetSkip((paginate - 1) * 20)

	cursor, err := col.Find(ctx, bson.M{}, option)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}
	for cursor.Next(context.TODO()) {
		var register models.Product
		err := cursor.Decode(&register)
		if err != nil {
			return results, false
		}
		results = append(results, &register)
	}
	return results, true
}
