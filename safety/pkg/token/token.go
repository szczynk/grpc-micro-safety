package token

import (
	"fmt"
	"safety/pkg/grpc_errors"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type Maker interface {
	CreateToken(userID uuid.UUID, userRole string, duration time.Duration) (string, paseto.JSONToken, error)
	VerifyToken(token string, expiration bool) (*paseto.JSONToken, error)
}

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (*PasetoMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	return &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *PasetoMaker) CreateToken(userID uuid.UUID, userRole string, duration time.Duration) (string, paseto.JSONToken, error) {
	sessionID, err := uuid.NewRandom()
	if err != nil {
		return "", paseto.JSONToken{}, err
	}

	now := time.Now()
	payload := paseto.JSONToken{
		Audience:   "szcz",
		Issuer:     "szcz_grpc_auth_svc",
		Jti:        "szcz_grpc",
		Subject:    "szcz_grpc_auth_sess",
		IssuedAt:   now,
		Expiration: now.Add(duration),
		NotBefore:  now,
	}
	payload.Set("session_id", sessionID.String())
	payload.Set("user_id", userID.String())
	payload.Set("role", userRole)

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, "Szczynk Initiative Enterprises")
	return token, payload, err
}

func (maker *PasetoMaker) VerifyToken(token string, expiration bool) (*paseto.JSONToken, error) {
	payload := new(paseto.JSONToken)
	var footer string

	if len(token) == 0 {
		return nil, grpc_errors.ErrNoRefreshToken
	}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, &payload, &footer)
	if err != nil {
		return nil, err
	}

	if expiration {
		err = payload.Validate(
			paseto.ForAudience("szcz"),
			paseto.IdentifiedBy("szcz_grpc"),
			paseto.IssuedBy("szcz_grpc_auth_svc"),
			paseto.Subject("szcz_grpc_auth_sess"),
			paseto.ValidAt(time.Now()),
		)
		if err != nil {
			return nil, err
		}
		return payload, nil
	} else {
		err = payload.Validate(
			paseto.ForAudience("szcz"),
			paseto.IdentifiedBy("szcz_grpc"),
			paseto.IssuedBy("szcz_grpc_auth_svc"),
			paseto.Subject("szcz_grpc_auth_sess"),
		)
		if err != nil {
			return nil, err
		}
		return payload, nil
	}
}
