package token

import (
	"fmt"
	"gin-base-api/pkg/config"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenService struct {
	config *config.Config
}

var tokenServiceInstance *TokenService

func NewService(config *config.Config) *TokenService {
	if tokenServiceInstance == nil {
		tokenServiceInstance = &TokenService{
			config: config,
		}
	}
	return tokenServiceInstance
}

func (tokenService *TokenService) NewAccessToken(claims *UserClaims) (*AccessToken, error) {
	expirationTime := tokenService.config.Jwt.ExpirationTime
	claims.ExpiresAt = time.Now().Add(time.Duration(expirationTime) * time.Second).Unix()

	unsignedAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := unsignedAccessToken.SignedString([]byte(tokenService.config.Jwt.Secret))

	if err != nil {
		return nil, err
	}

	return &AccessToken{
		Token:          accessToken,
		ExpirationTime: expirationTime,
	}, nil
}

func (tokenService *TokenService) ParseAccessToken(accessToken string) (*UserClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenService.config.Jwt.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	userClaims, ok := parsedAccessToken.Claims.(*UserClaims)

	if !ok {
		return nil, fmt.Errorf("error while parsing claims")
	}

	return userClaims, nil

}

func (tokenService *TokenService) IsAuthorized(accessToken string, secret string) (bool, error) {
	_, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
