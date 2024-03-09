package controllers

import (
	"time"

	"github.com/golang-jwt/jwt"
)


func generateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	})
	
	return token.SignedString(jwtKey)
}