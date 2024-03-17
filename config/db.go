package config

import (
	"fmt"
	"os"
	"time"
)

type DBConfig struct {
	PSQLConfig PSQLConfig
	LocalCache LocalCacheConfig
}

var (

	// DB ENV VARS
	DB_HOST     = os.Getenv("DB_HOST")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME     = os.Getenv("DB_NAME")
	DB_PORT     = os.Getenv("DB_PORT")
)

func GetDBConfig() *DBConfig {
	dbUrl := ""
	if os.Getenv("ENV") == "production" {

		dbUrl = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=verify-full&sslrootcert=ap-southeast-1-bundle.pem",
			DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	} else {
		dbUrl = conf.GetOptionalValue("SHOPIFYX_BACKEND_DB_MASTER", "postgres://loco:loco@localhost:5432/shopifyx_db?sslmode=disable")
	}
	return &DBConfig{

		PSQLConfig: PSQLConfig{
			Master:        dbUrl,
			Slave:         dbUrl,
			LogLevel:      conf.GetOptionalIntValue("SHOPIFYX_BACKEND_DB_LOG_LEVEL", 4),
			MaxLifeTime:   GetTimeDuration("SHOPIFYX_BACKEND_DB_MAX_LIFETIME", 1*time.Hour),
			MaxConnection: conf.GetOptionalIntValue("SHOPIFYX_BACKEND_DB_MAX_CONNECTION", 50),
		},
		LocalCache: LocalCacheConfig{
			DefaultExpiration: GetTimeDuration("SHOPIFYX_BACKEND_LOCAL_CACHE_DEFAULT_EXPIRATION_TIME", 1*time.Minute),
			CleanupInterval:   GetTimeDuration("SHOPIFYX_BACKEND_LOCAL_CACHE_CLEANUP_INTERVAL", 1*time.Minute),
		},
	}
}

type LocalCacheConfig struct {
	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
}

type PSQLConfig struct {
	Master        string
	Slave         string
	LogLevel      int
	MaxLifeTime   time.Duration
	MaxConnection int
}
