package repository

import (
	"auth/internal/domain"
	"auth/internal/models"
	"auth/pkg/grpc_errors"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

const basePrefix string = "sessions:"

// Session repository
type sessionRepo struct {
	redisClient *redis.Client
	basePrefix  string
}

// Session repository constructor
func NewSessionRepository(redisClient *redis.Client) domain.SessionRepository {
	return &sessionRepo{redisClient: redisClient, basePrefix: basePrefix}
}

// *Command

// Create session in redis
func (s *sessionRepo) CreateSession(ctx context.Context, maxAge time.Duration, session *models.Session) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "SessionRepo.CreateSession")
	defer span.Finish()

	sessionKey := s.createKey(session.ID.String())

	sessionBytes, err := json.Marshal(&session)
	if err != nil {
		return err
	}

	return s.redisClient.Set(ctx, sessionKey, sessionBytes, maxAge).Err()
}

// Delete session by id
func (s *sessionRepo) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "SessionRepo.DeleteByID")
	defer span.Finish()

	return s.redisClient.Del(ctx, s.createKey(ID.String())).Err()
}

// *Query

// Get session by id
func (s *sessionRepo) FindByID(ctx context.Context, ID uuid.UUID) (*models.Session, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "SessionRepo.FindByID")
	defer span.Finish()

	sessionBytes, err := s.redisClient.Get(ctx, s.createKey(ID.String())).Bytes()
	if err != nil {
		if err != redis.Nil {
			return nil, grpc_errors.ErrNotFound
		}
		return nil, err
	}

	session := new(models.Session)
	if err = json.Unmarshal(sessionBytes, &session); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *sessionRepo) createKey(sessionID string) string {
	return fmt.Sprintf("%s: %s", s.basePrefix, sessionID)
}
