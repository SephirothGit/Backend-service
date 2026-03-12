package config

import (
	"os"
	"strconv"
	"time"
)

type config struct {
	Port         string
	JWTSecret    string
	RateLimit    int
	TimeDuration time.Duration
}

func Load() Config {

	ratelimit, _ := strconv.Atoi(os.Getenv("RATE_LIMIT"))
	timeout, _ := strconv.Atoi(os.Getenv("REQUEST_TIMEOUT"))

	return Config{
		Port:           os.Getenv("PORT"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		RateLimit:      ratelimit,
		RequestTimeout: tume.Duration(timeout) * time.Second,
	}
}
