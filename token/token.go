package token

import "time"

type Token interface {
	CreateToken(data string, duration time.Duration) (string, error)
	ValidateToken(token string) (*Payload, error)
}
