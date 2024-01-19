package token

import "time"

type Token interface {
	CreateToken(string, time.Duration) (string, error)
	ValidateToken(string) (*Payload, error)
}
