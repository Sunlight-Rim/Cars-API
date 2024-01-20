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
	jwtClaims := make(jwt.MapClaims)

	if _, err := jwt.ParseWithClaims(token, jwtClaims, func(*jwt.Token) (any, error) {
		return s.secret, nil
	}); err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.Wrapf(errors.ExpiredToken, "expired token, %v", err)
		}
		return nil, errors.Wrapf(errors.InvalidToken, "invalid token, %v", err)
	}

	claims, err := parse(jwtClaims)
	if err != nil {
		return nil, errors.Wrap(err, "parse claims")
	}

	return claims, nil
}

// Parse expired token.
func (s *service) ParseExpired(token string) (*auth.Claims, error) {
	jwtClaims := make(jwt.MapClaims)

	if _, err := jwt.ParseWithClaims(token, jwtClaims, func(*jwt.Token) (any, error) {
		return s.secret, nil
	}); err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		return nil, errors.Wrapf(errors.InvalidToken, "invalid token, %v", err)
	}

	claims, err := parse(jwtClaims)
	if err != nil {
		return nil, errors.Wrap(err, "parse claims")
	}

	return claims, nil
}

// StoreUserRefresh saves refresh token by user ID to redis.
func (s *service) StoreUserRefresh(token string, userID uint64) error {
	if err := s.redis.Set(
		context.TODO(),
		fmt.Sprintf("%s_%d_%s", refreshKeyPrefix, userID, token),
		nil,
		s.refreshExp,
	).Err(); err != nil {
		return errors.Wrap(err, "save refresh token")
	}

	return nil
}

// RevokeUserRefresh deletes refresh token by user ID from redis.
func (s *service) RevokeUserRefresh(token string, userID uint64) error {
	if err := s.redis.GetDel(
		context.TODO(),
		fmt.Sprintf("%s_%d_%s", refreshKeyPrefix, userID, token),
	).Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return errors.Wrap(errors.InvalidToken, "token not in redis")
		}
		return errors.Wrap(err, "pop refresh token")
	}

	return nil
}

// RevokeUserAllRefresh deletes all refresh tokens by user ID from redis.
func (s *service) RevokeUserAllRefresh(userID uint64) error {
	var tokens []string

	if err := s.redis.Keys(
		context.TODO(),
		fmt.Sprintf("%s_%d_*", refreshKeyPrefix, userID),
	).ScanSlice(&tokens); err != nil {
		return errors.Wrap(err, "get all refresh tokens")
	}

	if err := s.redis.Del(
		context.TODO(),
		tokens...,
	).Err(); err != nil {
		return errors.Wrap(err, "remove all refresh tokens")
	}

	return nil
}

// Parse token claims to struct.
func parse(jwtClaims jwt.MapClaims) (*auth.Claims, error) {
	var claims auth.Claims

	// Numeric values is float64 for "encoding/json"
	userID, ok := jwtClaims["user_id"].(float64)
	if !ok {
		return nil, errors.Wrap(errors.InvalidToken, "user_id")
	}

	claims.UserID = uint64(userID)

	return &claims, nil
}
