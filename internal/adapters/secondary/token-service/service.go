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
func (s *service) Create(userID uint64) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(s.accessExp).Unix(),
		"user_id": userID,
	}).SignedString(s.secret)
	if err != nil {
		return "", errors.Wrap(err, "signing token")
	}

	return token, nil
}

// Parse token.
func (s *service) Parse(token string) (*auth.Claims, error) {
	var claims auth.Claims

	if _, err := jwt.ParseWithClaims(token, &claims, func(*jwt.Token) (any, error) {
		return s.secret, nil
	}); err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.Wrapf(errors.ExpiredToken, "expired token, %v", err)
		}
		return nil, errors.Wrapf(errors.InvalidToken, "invalid token, %v", err)
	}

	return &claims, nil
}

// Parse expired token.
func (s *service) ParseExpired(token string) (*auth.Claims, error) {
	var claims auth.Claims

	if _, err := jwt.ParseWithClaims(token, &claims, func(*jwt.Token) (any, error) {
		return s.secret, nil
	}); err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		return nil, errors.Wrapf(errors.InvalidToken, "invalid token, %v", err)
	}

	return &claims, nil
}

// StoreUserRefresh saves refresh token by user ID to redis.
func (s *service) StoreUserRefresh(userID uint64, token string) error {
	if err := s.redis.Set(
		context.TODO(),
		fmt.Sprintf("%s_%d_%s", refreshKeyPrefix, userID, token),
		nil,
		s.refreshExp,
	).Err(); err != nil {
		return errors.Wrap(err, "store token")
	}

	return nil
}

func (s *service) ListUserRefresh(userID uint64) ([]string, error) {
	var tokens []string

	if err := s.redis.Keys(
		context.TODO(),
		fmt.Sprintf("%s_%d_*", refreshKeyPrefix, userID),
	).ScanSlice(&tokens); err != nil {
		return nil, errors.Wrap(err, "get all tokens")
	}

	return tokens, nil
}

// RevokeUserRefresh deletes refresh token by user ID from redis.
func (s *service) RevokeUserRefresh(userID uint64, token string) error {
	if err := s.redis.GetDel(
		context.TODO(),
		fmt.Sprintf("%s_%d_%s", refreshKeyPrefix, userID, token),
	).Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return errors.Wrap(errors.InvalidToken, "token not in redis")
		}
		return errors.Wrap(err, "revoke token")
	}

	return nil
}

// RevokeUserRefreshAll deletes all refresh tokens by user ID from redis.
func (s *service) RevokeUserRefreshAll(userID uint64) ([]string, error) {
	var tokens []string

	if err := s.redis.Keys(
		context.TODO(),
		fmt.Sprintf("%s_%d_*", refreshKeyPrefix, userID),
	).ScanSlice(&tokens); err != nil {
		return nil, errors.Wrap(err, "get all tokens")
	}

	if err := s.redis.Del(
		context.TODO(),
		tokens...,
	).Err(); err != nil {
		return nil, errors.Wrap(err, "revoke all tokens")
	}

	return tokens, nil
}
