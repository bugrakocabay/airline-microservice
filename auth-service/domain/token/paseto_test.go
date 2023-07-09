package token

import (
	"crypto/rand"
	"testing"
	"time"

	"github.com/aead/chacha20poly1305"
)

func TestNewPasetoMaker(t *testing.T) {
	key := make([]byte, chacha20poly1305.KeySize)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatalf("failed to generate random key: %v", err)
	}

	maker, err := NewPasetoMaker(string(key))
	if err != nil {
		t.Fatalf("failed to create PasetoMaker: %v", err)
	}

	if maker == nil {
		t.Error("maker is nil")
	}
}

func TestCreateToken(t *testing.T) {
	key := "12345678901234567890123456789012"

	maker, err := NewPasetoMaker(key)
	if err != nil {
		t.Fatalf("failed to create PasetoMaker: %v", err)
	}

	userID := "testUserID"
	email := "test@example.com"
	duration := time.Minute

	token, err := maker.CreateToken(userID, email, duration)
	if err != nil {
		t.Fatalf("failed to create token: %v", err)
	}

	if len(token) == 0 {
		t.Error("failed to create token: token is empty")
	}
}

func TestVerifyToken(t *testing.T) {
	key := "12345678901234567890123456789012"

	maker, err := NewPasetoMaker(key)
	if err != nil {
		t.Fatalf("failed to create PasetoMaker: %v", err)
	}

	userID := "testUserID"
	email := "test@example.com"
	duration := time.Minute

	token, err := maker.CreateToken(userID, email, duration)
	if err != nil {
		t.Fatalf("failed to create token: %v", err)
	}

	payload, err := maker.VerifyToken(token)
	if err != nil {
		t.Fatalf("failed to verify token: %v", err)
	}

	if payload.UserID != userID {
		t.Errorf("payload.UserID does not match: got %v, want %v", payload.UserID, userID)
	}

	if payload.Email != email {
		t.Errorf("payload.Email does not match: got %v, want %v", payload.Email, email)
	}
}
