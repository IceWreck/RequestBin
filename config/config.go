package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//Config is the settings struct
type Config struct {
	Username string
	Password string
	Port     string
}

//LoadSettings loads your configuration from your dotenv
func LoadSettings() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	settings := Config{
		Username: os.Getenv("REQUESTBIN_USER"),
		Password: os.Getenv("REQUESTBIN_PASSWORD"),
		Port:     ":" + os.Getenv("REQUESTBIN_PORT"),
	}

	return settings

}
