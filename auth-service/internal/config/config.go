package config

import (
	"log"
	"strconv"

	osutils "github.com/krissolui/go-utils/os-utils"
)

const (
	defaultPort              = "80"
	defaultAttempts          = 10
	defaultDelayInSecond     = 2
	defaultSessionServiceURL = "http://session-service"
)

type StorageConfig struct {
	DSN           string
	Attempts      int
	DelayInSecond int
}

type Config struct {
	Port              string
	Storage           StorageConfig
	SessionServiceURL string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) LoadEnv() {
	c.Port = osutils.GetEnv("PORT", defaultPort)
	c.SessionServiceURL = osutils.GetEnv("SESSION_SERVICE_URL", defaultSessionServiceURL)

	dsn := osutils.GetEnv("DSN")
	if dsn == "" {
		log.Fatal("DSN is required but not found!")
	}
	attempts, err := strconv.Atoi(osutils.GetEnv("DB_ATTEMPTS"))
	if err != nil {
		attempts = defaultAttempts
	}
	delayInSecond, err := strconv.Atoi(osutils.GetEnv("DB_DELAY_IN_SECOND"))
	if err != nil {
		delayInSecond = defaultDelayInSecond
	}

	c.Storage = StorageConfig{
		DSN:           dsn,
		Attempts:      attempts,
		DelayInSecond: delayInSecond,
	}
}
