package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTToken struct {
	method    jwt.SigningMethod
	secretKey string
}

func (t JWTToken) CreateToken(data string, duration time.Duration) (string, error) {
	payload := NewPayload(data, time.Now().Add(duration))
	token := jwt.NewWithClaims(t.method, payload)
	return token.SignedString([]byte(t.secretKey))
}

func (t JWTToken) ValidateToken(token string) (*Payload, error) {
	payload := &Payload{}
	_, err := jwt.ParseWithClaims(token, payload, func(token *jwt.Token) (interface{}, error) {
		if token.Method != t.method {
			return nil, ErrInvalidToken
		}
		return []byte(t.secretKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	return payload, nil
}

func NewJWTToken(secretKey string, method jwt.SigningMethod) Token {
	return &JWTToken{
		secretKey: secretKey,
		method:    method,
	}
}
