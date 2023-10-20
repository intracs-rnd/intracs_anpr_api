package auth

import (
	"errors"
	"fmt"
	"intracs_anpr_api/internal/env"
	"intracs_anpr_api/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user model.User) model.AuthToken {
	secret := env.Get("APP_KEY")
	expiresAt := time.Now().Add(time.Minute * 15).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = &model.AuthTokenClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		User: user,
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("Token generate failed.", err)
	}

	return model.AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	}
}

func VerifyJWT(tokenString string) (model.User, error) {
	secret := env.Get("APP_KEY")
	claims := jwt.MapClaims{}
	user := model.User{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return user, errors.New("signature invalid")
		}

		return user, errors.New("could not parse the auth token")
	}

	if !token.Valid {
		return user, errors.New("invalid token")
	}
	fmt.Println("TOKEN is :", token.Valid)

	decoded := make(map[string]interface{})
	for key, val := range claims {
		decoded[key] = val
	}

	var email string
	if keyExists(decoded, "email") {
		email = decoded["email"].(string)
		user.Email = email
	}

	var createdAt time.Time
	if keyExists(decoded, "created_at") {
		createdAt = decoded["created_at"].(time.Time)
		user.CreatedAt = createdAt
	}

	return user, nil
}

func keyExists(decoded map[string]interface{}, key string) bool {
	val, ok := decoded[key]
	return ok && val != nil
}
