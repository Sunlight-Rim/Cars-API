package server

import (
	"context"
	"fmt"
	"net/http"

	"cars/internal/adapters/primary/graphql"
	"cars/internal/adapters/primary/rest"
	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type server struct {
	echo    *echo.Echo
	rest    *rest.Service
	graphql *graphql.Service
}

func New(rest *rest.Service, graphql *graphql.Service) *server {
	s := server{
		echo:    echo.New(),
		rest:    rest,
		graphql: graphql,
	}

	s.echo.Use(echomw.Recover(), echomw.RequestID(), LoggerMW())
	s.setRoutes()

	// Custom Echo error handler
	s.echo.HTTPErrorHandler = func(err error, c echo.Context) {
		if httpError, ok := err.(*echo.HTTPError); ok {
			switch httpError.Code {
			case http.StatusNotFound:
				c.JSONBlob(mapResponse(nil, errors.NotFound))

			case http.StatusMethodNotAllowed:
				c.JSONBlob(mapResponse(nil, errors.MethodNotAllowed))
			}
		}
	}

	return &s
}

// Start starts the server.
func (s *server) Start() error {
	return s.echo.Start(fmt.Sprintf(
		"%s:%s",
		viper.GetString("server.service.host"),
		viper.GetString("server.service.port"),
	))
}

// StartPprof starts handling pprof profiler.
func (s *server) StartPprof() error {
	e := echo.New()
	e.HideBanner = true

	e.GET("/debug/*", echo.WrapHandler(http.DefaultServeMux))

	return e.Start(fmt.Sprintf(
		"%s:%s",
		viper.GetString("server.pprof.host"),
		viper.GetString("server.pprof.port"),
	))
}

// Shutdown provides server gracefully shutdown.
func (s *server) Shutdown(ctx context.Context) error {
	if err := s.echo.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "echo shutdown")
	}

	return nil
}
