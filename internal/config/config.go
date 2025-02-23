package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


type ApplicationConfig struct {
	DbConfig DBConfig
	RedisConfig RedisConfig
	Config Config
}

type Config struct {
	Port int
	SecretKey string
}


func LoadConfig() (*ApplicationConfig, error) {
	_ = godotenv.Load()
	dbConfig :=  DBConfig{
		Host: os.Getenv("DB_HOST"),
		Port: func() int {
			port, err := strconv.Atoi(os.Getenv("DB_PORT"))
			if err != nil {
				return 0
			}
			return port
		}(),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Driver: os.Getenv("DB_DRIVER"),
		SkipMigrations: func() bool {
			skip, err := strconv.ParseBool(os.Getenv("SKIP_MIGRATIONS"))
			if err != nil {
				return false
			}
			return skip
		}(),
		SkipSeeding: func() bool {
			skip, err := strconv.ParseBool(os.Getenv("SKIP_SEEDING"))
			if err != nil {
				return false
			}
			return skip
		}(),
	}

	appConfig := Config{
		Port: func() int {
			port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
			if err != nil {
				return 0
			}
			return port
		}(),
		SecretKey: os.Getenv("SECRET_KEY"),
	}

	redisConfig := RedisConfig{
		Host: os.Getenv("REDIS_HOST"),
		Port: func() int {
			port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
			if err != nil {
				return 0
			}
			return port
		}(),
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: func() int {
			db, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
			if err != nil {
				return 0
			}
			return db
		}(),
	}
	cfg := &ApplicationConfig{
		DbConfig: dbConfig,
		Config: appConfig,
		RedisConfig: redisConfig,
	}

	return cfg, nil;
}