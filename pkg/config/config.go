package config

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBDatabase string `mapstructure:"DB_DATABASE"`
	DBUrl      string `mapstructure:"DB_URL"`
	HOST       string `mapstructure:"HOST"`
	Address    string `mapstructure:"ADDRESS"`
	Debug      bool   `mapstructure:"DEBUG"`
}

var ConfGlob Config

func LoadConfig() (config Config, err error) {
	godotenv.Load(".env")
	path := "."

	if os.Getenv("ENV") == "dev" {

		viper.AddConfigPath(path)
		viper.SetConfigName("dev")
		viper.SetConfigType("env")

		viper.AutomaticEnv()

		err = viper.ReadInConfig()
		if err != nil {
			return
		}

		err = viper.Unmarshal(&config)

		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		log.Info().Msg("Load Development Environment...")
		return
	}
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	log.Info().Msg("Load Production Environment...")
	return
}
