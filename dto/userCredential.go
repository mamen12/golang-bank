package dto

import "github.com/golang-jwt/jwt"

type UserCredential struct {
	Id       string
	Username string
	jwt.StandardClaims
}
