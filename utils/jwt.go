package utils

import (
	
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = []byte("mysecretkey")

type Claims struct{
	UserID string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string, role string)(string, error)  {

	exp := time.Now().Add(24*time.Hour)

	claims:= &Claims{
		UserID: userID,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	return token.SignedString(SECRET_KEY)
}