package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
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

func ExtractTokenFromRequest(r *http.Request) (string, error) {
	tokenString, err := request.AuthorizationHeaderExtractor.ExtractToken(r)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUsertIdFromClaims(c *gin.Context) (string, error) {
	claims, exists := c.Get("claims")

	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ERROR:": "BAD_REQUEST"})
		return "", err
	}

	mapClaims := claims.(*Claims)

	return mapClaims.ID, nil
}
