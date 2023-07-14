package handlers

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/now"
	"github.com/joho/godotenv"
)

type Claims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

var err = godotenv.Load()
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(id string) (string, error) {
	if err != nil {
		panic("Error loading .env file")
	}

	claims := Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.BeginningOfDay().AddDate(0, 0, 7)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*Claims, error) {
	if err != nil {
		panic("Error loading .env file")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid token.")
	}

	return claims, nil
}
