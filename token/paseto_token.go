package token

import (
	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
	"time"
)

type PasetoToken struct {
	token     paseto.Token
	secretKey string
}

func (t PasetoToken) CreateToken(data string, duration time.Duration) (string, error) {
	expiredTime := time.Now().Add(duration)
	payload := NewPayload(data, expiredTime)
	t.token.SetExpiration(expiredTime)
	t.token.SetString("id", payload.ID.String())
	t.token.SetString("data", payload.Data)
	key, err := paseto.V4SymmetricKeyFromBytes([]byte(t.secretKey))
	if err != nil {
		return "", err
	}
	en := t.token.V4Encrypt(key, nil)
	return en, nil
}

func (t PasetoToken) ValidateToken(token string) (*Payload, error) {
	parser := paseto.NewParser()
	parser.AddRule(paseto.NotExpired())
	key, err := paseto.V4SymmetricKeyFromBytes([]byte(t.secretKey))
	if err != nil {
		return nil, err
	}
	parsedToken, err := parser.ParseV4Local(key, token, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	var payload Payload
	id, err := parsedToken.GetString("id")
	if err != nil {
		return nil, err
	}
	payload.ID, err = uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	payload.Data, err = parsedToken.GetString("data")
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func NewPasetoToken(secretKey string) Token {
	return &PasetoToken{
		secretKey: secretKey,
		token:     paseto.NewToken(),
	}
}
