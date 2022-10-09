package utils

import (
	// "crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

const SALT = 13

func GenerateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)
	// _, err := rand.Read(salt[:])

	// if err != nil {
	// 	panic(err)
	// }

	return salt
}

func HashPassword(password string, salt []byte) string {
	var passwordBytes = []byte(password)
	var sha512Hasher = sha512.New()
	passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)
	var hashedPasswordBytes = sha512Hasher.Sum(nil)
	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}

// VALIDAR COINCIDENCIA
func DoPasswordsMatch(hashedPassword, currPassword string,
	salt []byte) bool {
	var currPasswordHash = HashPassword(currPassword, salt)
	return hashedPassword == currPasswordHash
}

func TestEncrypt() {
	var salt = GenerateRandomSalt(SALT)

	// Hash password using the salt
	var hashedPassword = HashPassword("password", salt)

	fmt.Println("Password Hash:", hashedPassword)
	fmt.Println("Salt:", salt)

	// Check if passed password matches the original password by hashing it
	// with the original password's salt and check if the hashes match
	fmt.Println("Password Match:", DoPasswordsMatch(hashedPassword, "password", salt))
}
