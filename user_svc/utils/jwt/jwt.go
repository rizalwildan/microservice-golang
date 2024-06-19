package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"microsrvc/user_svc/config"
	"time"
)

type TokenPayload struct {
	ID uint
}

func GenerateToken(payload *TokenPayload) string {
	v, err := time.ParseDuration(config.TOKEN_EXP)

	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(v).Unix(),
		"ID":  payload.ID,
	})

	token, err := t.SignedString([]byte(config.TOKEN_KEY))

	if err != nil {
		panic(err)
	}

	return token
}

func parse(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.TOKEN_KEY), nil
	})
}

func VerifyToken(token string) (*TokenPayload, error) {
	parsed, err := parse(token)

	if err != nil {
		return nil, err
	}

	// parsing token claims
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	// Getting ID
	id, ok := claims["ID"].(float64)
	if !ok {
		return nil, errors.New("something wrong")
	}

	return &TokenPayload{ID: uint(id)}, nil
}
