package db

import (
	"product-service/config" // Importa el paquete de configuración local
	"product-service/models" // Importa los modelos de datos locales

	"gorm.io/driver/postgres" // Driver de PostgreSQL para GORM
	"gorm.io/gorm"            // ORM (Object-Relational Mapping) para Go
)

// DB es una variable global que almacena la conexión a la base de datos
var DB *gorm.DB

// ConnectDatabase establece la conexión con la base de datos
func ConnectDatabase(config *config.Config) {

	// Abre una conexión a la base de datos PostgreSQL usando la cadena de conexión proporcionada
	database, err := gorm.Open(postgres.Open(config.DBConnectionString), &gorm.Config{})
	if err != nil {
		// Si hay un error al conectar, detiene la ejecución del programa
		panic("Error al conectar a la base de datos")
	}

	// AutoMigrate crea o actualiza las tablas en la base de datos según los modelos definidos
	// En este caso, crea/actualiza la tabla para el modelo Product
	database.AutoMigrate(&models.Product{})

	// Asigna la conexión establecida a la variable global DB
	DB = database
}
