package config

import (
	"encoding/json"
	"log"
	"os"
)

func GetAppConfig() *AppConfig {

	return &AppConfig{
		JwtConfig:        GetJwtConfig(),
		ServerConfig:     GetServerConfig(),
		TranslatorConfig: GetTranslatorConfig(),
	}

}

type TranslatorConfig struct {
	ErrorFilePath string `json:"error_file_path"`
}

type AppConfig struct {
	JwtConfig        JwtConfig
	ServerConfig     ServerConfig
	TranslatorConfig TranslatorConfig
	DBConfig         DBConfig
}

type JwtConfig struct {
	Secret    string
	PinSecret string
}

type ServerConfig struct {
	Environment string
}

func GetJwtConfig() JwtConfig {
	secret := ""
	if os.Getenv("ENV") == "production" {
		secret = os.Getenv("JWT_SECRET")
	} else {
		secret = conf.GetOptionalValue("JWT_SECRET", "secret")
	}

	return JwtConfig{
		Secret:    secret,
		PinSecret: secret,
	}

}

func GetServerConfig() ServerConfig {
	return ServerConfig{
		Environment: conf.GetOptionalValue("APP_ENVIRONMENT", "Development"),
	}
}

func GetTranslatorConfig() TranslatorConfig {
	var cfg TranslatorConfig
	// err := json.Unmarshal([]byte(conf.GetValue("TRANSLATOR_CONFIG")), &cfg)
	err := json.Unmarshal([]byte("{}"), &cfg)
	if err != nil {
		log.Panicln("failed to unmarshal translator config:", err)
	}
	return cfg
}
