package models

import "github.com/golang-jwt/jwt"

// Claims model
type AppClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
