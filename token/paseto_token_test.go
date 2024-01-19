package token

import (
	"github.com/dora-exku/go-util/random"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var ()

func TestCreatePasetoToken(t *testing.T) {
	key := random.Letter(32)
	data := "data"

	pasetoToken := NewPasetoToken(key)
	token, err := pasetoToken.CreateToken(data, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := pasetoToken.ValidateToken(token)
	require.NoError(t, err)
	require.Equal(t, data, payload.Data)
}

func TestInvalidPasetoToken(t *testing.T) {
	secretKey := random.Letter(32)
	_token := NewPasetoToken(secretKey)

	data := "data"
	duration := -time.Minute

	token, err := _token.CreateToken(data, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := _token.ValidateToken(token)
	require.Error(t, err)
	require.Nil(t, payload)
	require.Equal(t, ErrInvalidToken, err)
}
