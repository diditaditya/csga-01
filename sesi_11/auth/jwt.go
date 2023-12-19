package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var KEY = os.Getenv("JWT_SECRET")

type Claims struct {
	Username string
	jwt.RegisteredClaims
}

func CreateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(1440 * time.Minute)

	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	key := []byte(KEY)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func VerifyJWT(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}
		if method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid signing method")
		}

		return []byte(KEY), nil
	})

	if err != nil {
		return "", errors.New("Unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("Unauthorized")
	}

	return fmt.Sprintf("%s", claims["Username"]), nil
}
