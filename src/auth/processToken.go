package auth

import (
	"os"

	"errors"

	"github.com/ArthurQR98/e-commerce/src/models"
	"github.com/ArthurQR98/e-commerce/src/services"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var IDCustomer string

func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	jwtKey := os.Getenv("JWTKEY")
	myKey := []byte(jwtKey)
	claims := &models.Claim{}

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		_, find, _ := services.CheckIfExistUser(claims.Email)
		if find {
			Email = claims.Email
			IDCustomer = claims.ID.Hex()
		}
		return claims, find, IDCustomer, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalid")
	}
	return claims, false, string(""), err
}
