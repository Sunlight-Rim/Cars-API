package tokenService

import "github.com/redis/go-redis/v9"

type service struct {
	redis *redis.Client
}

func New(redis *redis.Client) *service {
	return &service{
		redis: redis,
	}
}

func (s *service) Create(claims map[string]string) (token string, err error)
func (s *service) Parse(token string) (claims map[string]string, err error)
func (s *service) ParseExpired(token string) (claims map[string]string, err error)
func (s *service) Save(token string) error
func (s *service) Delete(token string) error
