package server

import (
	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

// Response logger
func LoggerMW() func(next echo.HandlerFunc) echo.HandlerFunc {
	return echomw.RequestLoggerWithConfig(echomw.RequestLoggerConfig{
		LogURI:       true,
		LogMethod:    true,
		LogStatus:    true,
		LogLatency:   true,
		LogRemoteIP:  true,
		LogRequestID: true,
		LogUserAgent: true,
		LogError:     true,
		LogHeaders:   []string{"Authorization"},

		LogValuesFunc: func(c echo.Context, values echomw.RequestLoggerValues) error {
			errLocation, _ := errors.GetLocation(values.Error)

			log := logrus.WithFields(logrus.Fields{
				"uri":            values.URI,
				"method":         values.Method,
				"status":         values.Status,
				"latency":        values.Latency,
				"remote_ip":      values.RemoteIP,
				"request_id":     values.RequestID,
				"user_agent":     values.UserAgent,
				"headers":        values.Headers,
				"error":          values.Error,
				"error_location": errLocation,
			})

			switch _, ok := errors.GetCode(values.Error); {
			case values.Error == nil:
				log.Info("Successful")

			case ok:
				log.Warn("Catched error")

			case errors.Is(values.Error, echo.ErrNotFound),
				errors.Is(values.Error, echo.ErrMethodNotAllowed):
				log.Warn("Not found")

			default:
				log.Error("Unregistered internal error")
			}

			return nil
		},
	})
}
