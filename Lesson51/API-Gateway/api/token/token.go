package token

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	signingKey = "jwt-token-secret"
)

type Token struct {
	AccessToken string `json:"access_token"`
}

func GeneratedJWTToken(userId, email, username string) *Token {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)

	claims["user_id"] = userId
	claims["email"] = email
	claims["username"] = username
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()
	access, err := accessToken.SignedString([]byte(signingKey))
	if err != nil {
		log.Fatal(err)
	}

	return &Token{AccessToken: access}
}

func ValidateToken(tokenStr string) (bool, error) {
	_, err := ExtractClaim(tokenStr)

	if err != nil {
		return false, err
	}

	return true, nil
}

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) { return []byte(signingKey), nil })
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
