package services

import (
	"context"
	"log"
	"time"

	"github.com/ArthurQR98/e-commerce/config"
	"github.com/ArthurQR98/e-commerce/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetReviews(paginate int64) ([]*models.Review, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := config.MongoCN.Database("ecommerce")
	col := db.Collection("reviews")

	var results []*models.Review
	option := options.Find()
	option.SetLimit(20)
	option.SetSkip((paginate - 1) * 20)

	cursor, err := col.Find(ctx, bson.M{}, option)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}
	for cursor.Next(context.TODO()) {
		var register models.Review
		err := cursor.Decode(&register)
		if err != nil {
			return results, false
		}
		results = append(results, &register)
	}
	return results, true
}
