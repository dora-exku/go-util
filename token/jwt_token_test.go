package token

import (
	"github.com/dora-exku/go-util/random"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJWTToken(t *testing.T) {
	secretKey := random.Letter(32)
	signingMethod := jwt.SigningMethodHS256
	jwtToken := NewJWTToken(secretKey, signingMethod)

	data := "data"
	duration := time.Minute

	expiredAt := time.Now().Add(duration)

	token, err := jwtToken.CreateToken(data, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := jwtToken.ValidateToken(token)
	require.NoError(t, err)
	require.Equal(t, data, payload.Data)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt.Time, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	secretKey := random.Letter(32)
	signingMethod := jwt.SigningMethodHS256
	jwtToken := NewJWTToken(secretKey, signingMethod)

	data := "data"
	duration := -time.Minute

	token, err := jwtToken.CreateToken(data, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := jwtToken.ValidateToken(token)
	require.Error(t, err)
	require.Nil(t, payload)
	require.Equal(t, ErrExpiredToken, err)
}

func TestInvalidJWTToken(t *testing.T) {
	secretKey := random.Letter(32)
	signingMethod := jwt.SigningMethodHS256
	jwtToken := NewJWTToken(secretKey, signingMethod)

	data := "data"
	duration := time.Minute

	token, err := jwtToken.CreateToken(data, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	anotherSecretKey := random.Letter(32)
	anotherSigningMethod := jwt.SigningMethodHS384
	anotherJWTToken := NewJWTToken(anotherSecretKey, anotherSigningMethod)

	payload, err := anotherJWTToken.ValidateToken(token)
	require.Error(t, err)
	require.Nil(t, payload)
	require.Equal(t, ErrInvalidToken, err)
}
