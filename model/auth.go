package model

import "github.com/dgrijalva/jwt-go"

type AuthToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
	ExpiresIn int64  `json:"expires_in"`
}

type AuthTokenClaim struct {
	jwt.StandardClaims
	User
}
