package server

import (
	"context"
	"fmt"
	"net/http"

	"cars/internal/adapters/primary/graphql"
	"cars/internal/adapters/primary/rest"
	"cars/internal/adapters/primary/rest/response"
	"cars/internal/domain/auth"
	"cars/internal/domain/cars"
	"cars/internal/domain/users"
	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type server struct {
	echo *echo.Echo

	rest    *rest.Handlers
	graphql *graphql.Handlers
}

func New(auth auth.IUsecase, users users.IUsecase, cars cars.IUsecase) *server {
	s := server{
		echo: echo.New(),

		rest:    rest.New(auth, users, cars),
		graphql: graphql.New(auth, users, cars),
	}

	s.echo.Use(echomw.Recover(), echomw.CORS(), echomw.RequestID(), LoggerMW())

	// Custom Echo error handler
	s.echo.HTTPErrorHandler = func(err error, c echo.Context) {
		var httpError *echo.HTTPError

		if errors.As(err, &httpError) {
			switch httpError.Code {
			case http.StatusNotFound:
				c.JSONBlob(response.Map(nil, errors.NotFound))

			case http.StatusMethodNotAllowed:
				c.JSONBlob(response.Map(nil, errors.MethodNotAllowed))
			}
		}
	}

	return &s
}

// Start server.
func (s *server) Start() error {
	return s.echo.Start(fmt.Sprintf(
		"%s:%s",
		viper.GetString("server.service.host"),
		viper.GetString("server.service.port"),
	))
}

// Shutdown server.
func (s *server) Shutdown(ctx context.Context) error {
	healthStatus = []byte("Dying...")

	if err := s.echo.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "echo shutdown")
	}

	return nil
}
