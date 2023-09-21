package handlers

import (
	"errors"
	"github/ericoliveiras/basic-crud-go/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/now"
)

type Claims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

var jwtSecret = config.LoadAuthConfig().AccessSecret

func GenerateToken(id string) (string, error) {
	claims := Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.BeginningOfDay().AddDate(0, 0, 7)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid token.")
	}

	return claims, nil
}

func GetUsertIdFromClaims(ctx *gin.Context) (string, error) {
	claims, exists := ctx.Get("claims")

	if !exists {
		return "", errors.New("Failed to get user id from claims.")
	}

	mapClaims := claims.(*Claims)

	return mapClaims.ID, nil
}
