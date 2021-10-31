package services

import (
	"github.com/ArthurQR98/e-commerce/src/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginCustomer(email string, password string) (models.Customer, bool) {
	customer, find, _ := CheckIfExistUser(email)
	if !find {
		return customer, false
	}
	passwordByte := []byte(password)
	passwordDB := []byte(customer.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordByte)
	if err != nil {
		return customer, false
	}
	return customer, true
}
