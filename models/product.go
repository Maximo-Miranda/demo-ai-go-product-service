package models

import "gorm.io/gorm" // Importa el paquete gorm para el mapeo objeto-relacional

// Product representa la estructura de un producto en la base de datos
type Product struct {
	gorm.Model          // Incorpora los campos ID, CreatedAt, UpdatedAt, y DeletedAt de gorm
	Name        string  `json:"name"`        // Nombre del producto
	Description string  `json:"description"` // Descripción del producto
	Price       float64 `json:"price"`       // Precio del producto
	Quantity    int     `json:"quantity"`    // Cantidad disponible del producto
}
