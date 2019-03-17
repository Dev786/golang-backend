package password

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashedPassword(plainPassword []byte) string {
	hashedPassword, err := bcrypt.GenerateFromPassword(plainPassword, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hashedPassword)
}

func VerifyPassword(hashedPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return false
	}
	return true
}
