package utils

import (
	"context"
	"time"

	"github.com/ArthurQR98/e-commerce/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConnectDatabase(database string, collection string) (context.Context, *mongo.Collection, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	db := config.MongoCN.Database(database)
	col := db.Collection(collection)
	return ctx, col, cancel
}
