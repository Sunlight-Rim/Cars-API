package tokenService

import (
	"context"
	"fmt"
	"time"

	"cars/internal/domain/auth"
	"cars/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

const refreshKeyPrefix = "refresh"

type service struct {
	redis      *redis.Client
	secret     []byte
	accessExp  time.Duration
	refreshExp time.Duration
}

func New(redis *redis.Client, secret string, accessExp, refreshExp time.Duration) *service {
	return &service{
		redis:      redis,
		secret:     []byte(secret),
		accessExp:  accessExp * time.Minute,
		refreshExp: refreshExp * time.Minute,
	}
}

// Create new token.
func (s *service) Create(id uint64) (string, error) {
	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims(auth.Claims{
			"exp": time.Now().Add(s.accessExp).Unix(),
			"id":  id,
		}),
	).SignedString(s.secret)
	if err != nil {
		return "", errors.Wrap(err, "signing token")
	}

	return token, nil
}

// Parse token.
func (s *service) Parse(token string) (auth.Claims, error) {
	claims := jwt.MapClaims(auth.Claims{})

	if _, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (any, error) {
		return s.secret, nil
	}); err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.Wrapf(errors.ExpiredToken, "expired token, %v", err)
		}
		return nil, errors.Wrapf(errors.InvalidToken, "invalid token, %v", err)
	}

	return claims, nil
}

// Parse expired token.
func (s *service) ParseExpired(token string) (auth.Claims, error) {
	claims := jwt.MapClaims(auth.Claims{})

	if _, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (any, error) {
		return s.secret, nil
	}); err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		return nil, errors.Wrapf(errors.InvalidToken, "invalid token, %v", err)
	}

	return claims, nil
}

// SaveRefresh saves refresh token with user ID to redis.
func (s *service) SaveRefresh(token string, id uint64) error {
	if err := s.redis.Set(
		context.TODO(),
		fmt.Sprintf("%s_%d_%s", refreshKeyPrefix, id, token),
		nil,
		s.refreshExp,
	).Err(); err != nil {
		return errors.Wrap(err, "save refresh token")
	}

	return nil
}

// RemoveRefresh deletes refresh token with user ID from redis.
func (s *service) RemoveRefresh(token string, id uint64) error {
	if err := s.redis.GetDel(
		context.TODO(),
		fmt.Sprintf("%s_%d_%s", refreshKeyPrefix, id, token),
	).Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return errors.Wrap(errors.InvalidToken, "token not in redis")
		}
		return errors.Wrap(err, "pop refresh token")
	}

	return nil
}

// RemoveAllRefresh deletes all refresh tokens of user ID from redis.
func (s *service) RemoveAllRefresh(id uint64) error {
	if err := s.redis.Del(
		context.TODO(),
		fmt.Sprintf("%s_%d_*", refreshKeyPrefix, id),
	).Err(); err != nil {
		return errors.Wrap(err, "pop refresh token")
	}

	return nil
}
