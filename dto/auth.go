package dto

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
)

type JWTClaims struct {
	ID        uuid.UUID `json:"sub,omitempty"`
	ExpiresAt int64     `json:"exp,omitempty"`
	IssuedAt  int64     `json:"iat,omitempty"`
}

func (c JWTClaims) Valid(helper *jwt.ValidationHelper) (err error) {
	if helper.After(time.Unix(c.ExpiresAt, 0)) {
		err = errors.New("token has expired")
	}
	if helper.Before(time.Unix(c.IssuedAt, 0)) {
		err = errors.New("token used before issued")
	}
	return err
}
