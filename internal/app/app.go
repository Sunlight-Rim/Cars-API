package app

import (
	"context"
	"os/signal"
	"syscall"

	"cars/internal/adapters/primary/profiler"
	"cars/internal/adapters/primary/server"

	"github.com/sirupsen/logrus"
)

func Start() {
	readFlags()
	readConfig()
	initLogger()

	postgres := connectPostgres()
	redis := connectRedis()

	// Init server
	s := server.New(
		newAuth(postgres, redis),
		newUsers(postgres),
		newCars(postgres),
	)

	// Start server
	go func() {
		logrus.Fatalf("Server: %v", s.Start())
	}()

	go func() {
		logrus.Fatalf("Prof: %v", profiler.Start())
	}()

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	<-ctx.Done()
	logrus.Fatalf("Shutdown: %v", s.Shutdown(ctx))
}
