package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConnectionString string
	APPEnv             string
	APPUserServiceUrl  string
}

func (c *Config) LoadConfig() {

	if os.Getenv("APP_ENV") != "production" {
		// Cargar archivo .env si existe
		err := godotenv.Load()
		if err != nil {
			log.Println("No se encontró el archivo .env")
		}
	}

	// Leer variables de entorno
	c.DBConnectionString = os.Getenv("DB_CONNECTION_STRING")
	c.APPEnv = os.Getenv("APP_ENV")
	c.APPUserServiceUrl = os.Getenv("APP_USER_SERVICE_URL")
}
