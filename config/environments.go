package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConnectionString string
	APPEnv             string
}

func (c *Config) LoadConfig() {
	// Cargar archivo .env si existe
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No se encontró el archivo .env")
	}

	// Leer variables de entorno
	c.DBConnectionString = os.Getenv("DB_CONNECTION_STRING")
	c.APPEnv = os.Getenv("APP_ENV")
}
