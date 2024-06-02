package token

import "github.com/golang-jwt/jwt"

type UserClaims struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}
