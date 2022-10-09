package handlers

import (
	"fmt"
	"server/utils"
)

func CreateCli() {
	pass := "ultrasecret"
	passEncrypt := utils.HashPassword(pass, utils.GenerateRandomSalt(15))
	fmt.Println("passEncryp:", passEncrypt)
	// data :=
	// err = json.NewDecoder(req.Body).Decode(&data)
}

func ValidatePass() {
	hashed := "f9c19211b7f71f77492ae2a22ff84e4deb7cb1a9b5a2ef624472d8f4a5df50b3fd9d66ab7a753cbfa7c81d5f60544bc3ecfbb9e0a315bfc1b2c75919049d5537"
	match := utils.DoPasswordsMatch(hashed, "ultrasecret", utils.GenerateRandomSalt(15))
	fmt.Println("Exist match: ", match)

}
