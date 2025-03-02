package config

import (
	"os"
	"strconv"
)

// Config structure
type Config struct {
	Server struct {
		Env     string
		ApiPort string
	}

	Log struct {
		Path         string
		ApiName      string
		RotationTime int
		MaxAge       int
	}

	Redis struct {
		Host     string
		Port     string
		Password string
		Db       int
		Key      struct {
			CountryCity        string
			CluesMapping       string
			FunFactMapping     string
			TriviaMapping      string
			CountryClues       string
			CountryFunFact     string
			CountryTrivia      string
			CountryCityClues   string
			CountryCityFunfact string
			CountryCityTrivia  string
			User               string
		}
	}
}

// Global Config instance
var C Config

// Load config from environment variables
func LoadConfig() {
	C.Server.Env = getEnv("SERVER_ENV", "live")
	C.Server.ApiPort = getEnv("SERVER_API_PORT", ":5000")

	C.Log.Path = getEnv("LOG_PATH", "./logs/")
	C.Log.ApiName = getEnv("LOG_API_NAME", "affluenceAPILogs-%Y-%m-%d.log")
	C.Log.RotationTime = getEnvInt("LOG_ROTATION_TIME", 24)
	C.Log.MaxAge = getEnvInt("LOG_MAX_AGE", 0)

	C.Redis.Host = getEnv("REDIS_HOST", "redis-production-8b2c.up.railway.app")
	C.Redis.Port = getEnv("REDIS_PORT", "6379")
	C.Redis.Password = getEnv("REDIS_PASSWORD", "")
	C.Redis.Db = getEnvInt("REDIS_DB", 0)

	C.Redis.Key.CountryCity = getEnv("REDIS_KEY_COUNTRY_CITY", "_country_cities")
	C.Redis.Key.CluesMapping = getEnv("REDIS_KEY_CLUES_MAPPING", "_clues")
	C.Redis.Key.FunFactMapping = getEnv("REDIS_KEY_FUNFACT_MAPPING", "_funFacts")
	C.Redis.Key.TriviaMapping = getEnv("REDIS_KEY_TRIVIA_MAPPING", "_trivia")
	C.Redis.Key.CountryClues = getEnv("REDIS_KEY_COUNTRY_CLUES", "_country_clues")
	C.Redis.Key.CountryFunFact = getEnv("REDIS_KEY_COUNTRY_FUNFACT", "_country_fun_facts")
	C.Redis.Key.CountryTrivia = getEnv("REDIS_KEY_COUNTRY_TRIVIA", "_country_trivia")
	C.Redis.Key.CountryCityClues = getEnv("REDIS_KEY_COUNTRY_CITY_CLUES", "_country_city_clues")
	C.Redis.Key.CountryCityFunfact = getEnv("REDIS_KEY_COUNTRY_CITY_FUNFACT", "_country_city_fun_facts")
	C.Redis.Key.CountryCityTrivia = getEnv("REDIS_KEY_COUNTRY_CITY_TRIVIA", "_country_city_trivia")
	C.Redis.Key.User = getEnv("REDIS_KEY_USER", "_users")
}

// Helper functions
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		intValue, err := strconv.Atoi(value)
		if err == nil {
			return intValue
		}
	}
	return fallback
}
