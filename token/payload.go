package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
)

type Payload struct {
	ID        uuid.UUID        `json:"id"`
	Data      string           `json:"data"`
	ExpiredAt *jwt.NumericDate `json:"expired_at"`
}

func (p Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return p.ExpiredAt, nil
}

func (p Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Now()), nil
}

func (p Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Now()), nil
}

func (p Payload) GetIssuer() (string, error) {
	return "", nil
}

func (p Payload) GetSubject() (string, error) {
	return "", nil
}

func (p Payload) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil
}

func NewPayload(data string, time time.Time) *Payload {
	return &Payload{
		ID:        uuid.New(),
		Data:      data,
		ExpiredAt: jwt.NewNumericDate(time),
	}
}
