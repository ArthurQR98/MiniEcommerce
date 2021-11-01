package auth

import (
	"os"

	"errors"

	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/services"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Email string
var IDCustomer primitive.ObjectID

func ProcessToken(tk string) (*models.Claim, bool, primitive.ObjectID, error) {
	jwtKey := os.Getenv("JWTKEY")
	myKey := []byte(jwtKey)
	claims := &models.Claim{}

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, find, _ := services.CheckIfExistUser(claims.Email)
		if find {
			Email = claims.Email
			IDCustomer = claims.ID
		}
		return claims, find, IDCustomer, nil
	}
	if !tkn.Valid {
		return claims, false, primitive.NilObjectID, errors.New("token invalid")
	}
	return claims, false, primitive.NilObjectID, err
}
