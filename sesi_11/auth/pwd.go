package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plain string) (string, error) {
	fmt.Println("Plain Password: ", plain)
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(plain string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err != nil
}
