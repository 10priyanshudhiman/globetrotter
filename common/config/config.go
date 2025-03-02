package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config represents the configuration structure
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

var C Config

func ReadConfig(path string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("49", err)
		// log.Fatalln(err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		log.Println("54", err)
		os.Exit(1)
	}
}
