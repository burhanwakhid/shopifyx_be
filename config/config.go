package config

import (
	"time"

	"github.com/gojekfarm/goconfig"
)

type Config struct {
	goconfig.BaseConfig
}

var conf *Config

func Initialize() *Config {
	conf = &Config{}
	conf.Load()
	return conf
}

func AppName() string {
	return conf.GetOptionalValue("APP_NAME", "")
}

func AppPort() int {
	return conf.GetOptionalIntValue("APP_PORT", 8000)
}

func RequestTimeoutSecs() int {
	return conf.GetOptionalIntValue("REQUEST_TIMEOUT_SECS", 30)
}

func GetConfig() *Config {
	return conf
}

func GetTimeDuration(key string, defaultTime time.Duration) time.Duration {
	// dur, err := time.ParseDuration(conf.GetOptionalValue(key, ""))
	// if err != nil {
	return defaultTime
	// }

	// return dur
}
