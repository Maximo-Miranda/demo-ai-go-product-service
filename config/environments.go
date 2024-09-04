package config

import (
	"log" // Importa la librería para imprimir mensajes de error
	"os"  // Importa la librería para acceder a variables de entorno

	"github.com/joho/godotenv" // Importa la librería para cargar variables de entorno desde un archivo .env
)

// Config estructura que almacena la configuración de la aplicación
type Config struct {
	DBConnectionString string // Cadena de conexión a la base de datos
	APPEnv             string // Entorno de la aplicación (development, production, etc.)
	APPUserServiceUrl  string // URL del servicio de usuarios
}

// LoadConfig carga la configuración desde variables de entorno o archivo .env
func (c *Config) LoadConfig() {

	// Verifica si la aplicación no está en producción
	if os.Getenv("APP_ENV") != "production" {
		// Intenta cargar el archivo .env si existe
		err := godotenv.Load()
		if err != nil {
			log.Println("No se encontró el archivo .env")
		}
	}

	// Lee las variables de entorno y las asigna a los campos de la estructura Config
	c.DBConnectionString = os.Getenv("DB_CONNECTION_STRING")
	c.APPEnv = os.Getenv("APP_ENV")
	c.APPUserServiceUrl = os.Getenv("APP_USER_SERVICE_URL")
}
