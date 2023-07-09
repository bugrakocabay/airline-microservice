package token

import (
	"errors"
	"os"
	"time"

	"github.com/o1egl/paseto"
)

var ErrExpiredToken = errors.New("token has expired")

// VerifyToken checks if the given token is valid
func VerifyToken(token string) (*Payload, error) {
	var pasetoMaker *paseto.V2
	payload := &Payload{}

	err := pasetoMaker.Decrypt(token, []byte(os.Getenv("SYMMETRIC_KEY")), payload, nil)
	if err != nil {
		return nil, err
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// Payload contains the payload data of the token
type Payload struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
