package handlers

import (
	"net/http"
	"strings"

	"github.com/ChrisCodeX/REST-API-Go/models"
	"github.com/ChrisCodeX/REST-API-Go/server"
	"github.com/golang-jwt/jwt"
)

func GetTokenAuthorizationHeader(s server.Server, w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	// Get the token from Authorization header
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))

	// Check the validation of the token
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
