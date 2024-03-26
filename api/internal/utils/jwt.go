package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

const minSecretKeySize = 32

// Payload contain data of the token
type Payload struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Role      string
	IssuedAt  time.Time
	ExpiredAt time.Time
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrorExpiredToken
	}

	return nil
}

// NewPayload creates a new Payload
func NewPayload(userID uuid.UUID, role string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		ID:        tokenID,
		UserID:    userID,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, err
}

// CreateToken generate token with userId and role
func CreateToken(userID uuid.UUID, role string, duration time.Duration, jwtSecret string) (string, error) {
	payload, err := NewPayload(userID, role, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, payload)

	return jwtToken.SignedString([]byte(jwtSecret))
}

// VerifyToken verify token
func VerifyToken(token string, jwtSecret string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrorInvalidToken
		}

		return []byte(jwtSecret), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		var verr *jwt.ValidationError
		ok := errors.As(err, &verr)
		if ok && errors.Is(verr, ErrorExpiredToken) {
			return nil, ErrorExpiredToken
		}

		return nil, ErrorInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrorInvalidToken
	}

	return payload, nil
}
