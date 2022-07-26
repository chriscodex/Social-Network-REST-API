package models

import "github.com/golang-jwt/jwt"

type AppClains struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
