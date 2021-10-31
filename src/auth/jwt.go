package auth

import (
	"os"
	"time"

	"github.com/ArthurQR98/e-commerce/src/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(customer models.Customer) (string, error) {
	jwt_key := os.Getenv("JWTKEY")
	MyKey := []byte(jwt_key)
	payload := jwt.MapClaims{
		"name":     customer.First_name,
		"lastname": customer.Last_name,
		"email":    customer.Email,
		"_id":      customer.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(MyKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
