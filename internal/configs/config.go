package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type SportsSyncIntervals struct {
	Soccer   time.Duration
	Football time.Duration
	Baseball time.Duration
}

type Config struct {
	HTTPServerAddr string
	GRPCServerAddr string
	LogLevel       string
	BaseURL        string
	SportsSyncIntervals
}

func LoadConfig() *Config {
	// 1. Load environment variables
	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load(".localEnv")
		if err != nil {
			log.Fatalf("failed to load .localEnv file: %v", err)
		}
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("failed to load .env file: %v", err)
		}
	}

	// 2. Parse environment variables
	return &Config{
		HTTPServerAddr: getEnv("HTTP_SERVER_ADDR", ":8080"),
		GRPCServerAddr: getEnv("GRPC_SERVER_ADDR", ":60061"),
		LogLevel:       getEnv("LOG_LEVEL", "info"),
		BaseURL:        getEnv("BASE_URL", "http://localhost:8080/lines/"),
		SportsSyncIntervals: SportsSyncIntervals{
			Soccer:   parseDuration(getEnv("SOCCER_SYNC_INTERVAL", "3s")),
			Football: parseDuration(getEnv("FOOTBALL_SYNC_INTERVAL", "4s")),
			Baseball: parseDuration(getEnv("BASEBALL_SYNC_INTERVAL", "5s")),
		},
	}
}

func parseDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		log.Fatalf("failed to parse duration: %v", err)
	}
	return d
}

func getEnv(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}
