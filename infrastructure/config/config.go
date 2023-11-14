package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Env                string `mapstructure:"APP_ENV"`
	ServerPort         string `mapstructure:"SERVER_PORT"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBUser             string `mapstructure:"DB_USER"`
	DBPass             string `mapstructure:"DB_PASS"`
	DBName             string `mapstructure:"DB_NAME"`
	AccessTokenSecret  string `mapstructure:"ACCESS_TOKEN_SECRET"`
	AccessTokenExp     int64  `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenSecret string `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExp    int64  `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
}

func GetConfig() Config {
	config := Config{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	log.Printf("The App is running in %v env\n", config.Env)

	return config
}
