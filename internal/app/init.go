package app

import (
	"flag"
	"os"

	authRepository "cars/internal/adapters/secondary/auth-repository"
	carsRepository "cars/internal/adapters/secondary/cars-repository"
	tokenService "cars/internal/adapters/secondary/token-service"
	usersRepository "cars/internal/adapters/secondary/users-repository"
	"cars/internal/domain/auth"
	"cars/internal/domain/cars"
	"cars/internal/domain/users"
	"cars/pkg/postgres"
	"cars/pkg/redis"
	"database/sql"

	"github.com/labstack/echo/v4"
	goRedis "github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var flags struct {
	configPath string
}

// Read application flags.
func readFlags() {
	flags.configPath = *flag.String("config", "configs/config.json", "path to configuration file")

	flag.Parse()
}

// Read configuration file.
func readConfig() {
	viper.SetConfigFile(flags.configPath)

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("read config: %v", err)
	}
}

// Initialize a logger.
func initLogger() {
	logLevel, err := logrus.ParseLevel(viper.GetString("logger.level"))
	if err != nil {
		logrus.Fatalf("parse logger level: %v", err)
	}

	logrus.SetLevel(logLevel)
	logrus.SetOutput(os.Stdout)

	if viper.GetBool("logger.json") {
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:   "2006-01-02 15:04:05.000",
			DisableHTMLEscape: true,
		})
		return
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
	})
}

// Connect to Postgres.
func connectPostgres() *sql.DB {
	db, err := postgres.Connect(postgres.ConnectionOptions{
		Host:            viper.GetString("database.postgres.host"),
		Port:            viper.GetString("database.postgres.port"),
		User:            viper.GetString("database.postgres.user"),
		Password:        viper.GetString("database.postgres.password"),
		DBName:          viper.GetString("database.postgres.database"),
		SSLMode:         viper.GetString("database.postgres.sslmode"),
		MaxOpenConns:    viper.GetInt("database.postgres.max_open_conns"),
		MaxIdleConns:    viper.GetInt("database.postgres.max_idle_conns"),
		ConnMaxLifetime: viper.GetDuration("database.postgres.conn_max_lifetime"),
		ConnMaxIdleTime: viper.GetDuration("database.postgres.conn_max_idle_time"),
	})
	if err != nil {
		logrus.Fatalf("connect to postgres: %v", err)
	}

	return db
}

// Connect to Redis.
func connectRedis() *goRedis.Client {
	client, err := redis.Connect(
		viper.GetString("database.redis.host"),
		viper.GetString("database.redis.port"),
		viper.GetString("database.redis.password"),
		viper.GetInt("database.redis.database"),
	)
	if err != nil {
		logrus.Fatalf("connect to redis: %v", err)
	}

	return client
}

// New Auth usecase with adapters.
func newAuth(postgres *sql.DB, redis *goRedis.Client) *auth.Usecase {
	return auth.New(
		authRepository.New(postgres),
		tokenService.New(
			redis,
			viper.GetString("token.secret"),
			viper.GetDuration("token.access_exp_minutes"),
			viper.GetDuration("token.refresh_exp_minutes"),
		),
	)
}

// New Users usecase with adapters.
func newUsers(postgres *sql.DB) *users.Usecase {
	return users.New(usersRepository.New(postgres))
}

// New Cars usecase with adapters.
func newCars(postgres *sql.DB) *cars.Usecase {
	return cars.New(carsRepository.New(postgres))
}

func newTokenChecker() echo.MiddlewareFunc {
	return tokenService.CheckTokenMW(viper.GetString("token.secret"))
}
