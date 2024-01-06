package app

import (
	"context"
	"os/signal"
	"syscall"

	"cars/internal/adapters/primary/graphql"
	"cars/internal/adapters/primary/rest"
	"cars/internal/adapters/primary/server"
	carsRepository "cars/internal/adapters/secondary/cars-repo"
	tokenService "cars/internal/adapters/secondary/token-service"
	usersRepository "cars/internal/adapters/secondary/users-repo"
	carsUsecase "cars/internal/usecases/cars"
	usersUsecase "cars/internal/usecases/users"

	"github.com/sirupsen/logrus"
)

func Start() {
	readFlags()
	readConfig()
	initLogger()

	postgres := connectPostgres()
	redis := connectRedis()

	// Init services
	var (
		users = usersUsecase.New(
			usersRepository.New(postgres),
			tokenService.New(redis),
		)

		cars = carsUsecase.New(
			carsRepository.New(postgres),
		)
	)

	// Init server
	s := server.New(
		rest.New(users, cars),
		graphql.New(users, cars),
	)

	// Start server
	go func() {
		logrus.Fatalf("Server: %v", s.Start())
	}()

	// Start pprof
	go func() {
		logrus.Fatalf("Prof: %v", s.StartPprof())
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
