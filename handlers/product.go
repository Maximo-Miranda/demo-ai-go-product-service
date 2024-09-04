package handlers

import (
	"net/http"               // Proporciona constantes HTTP y tipos de respuesta
	"product-service/db"     // Importa el paquete de base de datos local
	"product-service/models" // Importa los modelos de datos locales

	"github.com/labstack/echo/v4" // Framework web Echo para Go
	"gorm.io/gorm"                // ORM para Go, usado para manejar errores específicos de la base de datos
)

// CreateProduct maneja la creación de un nuevo producto
func CreateProduct(c echo.Context) error {
	// Crea una nueva instancia de Product
	product := new(models.Product)

	// Vincula los datos de la solicitud al producto
	if err := c.Bind(product); err != nil {
		// Si hay un error al vincular los datos, devuelve un error 400
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos de solicitud inválidos"})
	}

	// Intenta crear el producto en la base de datos
	if err := db.DB.Create(product).Error; err != nil {
		// Si hay un error al crear el producto, devuelve un error 500
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al crear el producto"})
	}

	// Devuelve el producto creado con código 201
	return c.JSON(http.StatusCreated, product)
}

// GetProduct maneja la obtención de un producto por su ID
func GetProduct(c echo.Context) error {
	// Obtiene el ID del producto de los parámetros de la URL
	id := c.Param("id")
	var product models.Product

	// Intenta obtener el producto de la base de datos
	if err := db.DB.First(&product, id).Error; err != nil {
		// Si el producto no se encuentra, devuelve un error 404
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Producto no encontrado"})
		}
		// Para otros errores, devuelve un error 500
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener el producto"})
	}

	// Devuelve el producto encontrado
	return c.JSON(http.StatusOK, product)
}

// UpdateProduct maneja la actualización de un producto existente
func UpdateProduct(c echo.Context) error {
	// Obtiene el ID del producto de los parámetros de la URL
	id := c.Param("id")
	var product models.Product

	// Intenta obtener el producto existente de la base de datos
	if err := db.DB.First(&product, id).Error; err != nil {
		// Si el producto no se encuentra, devuelve un error 404
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Producto no encontrado"})
		}
		// Para otros errores, devuelve un error 500
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener el producto"})
	}

	// Vincula los nuevos datos de la solicitud al producto
	if err := c.Bind(&product); err != nil {
		// Si hay un error al vincular los datos, devuelve un error 400
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos de solicitud inválidos"})
	}

	// Intenta guardar los cambios del producto en la base de datos
	if err := db.DB.Save(&product).Error; err != nil {
		// Si hay un error al actualizar, devuelve un error 500
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar el producto"})
	}

	// Devuelve el producto actualizado
	return c.JSON(http.StatusOK, product)
}

// DeleteProduct maneja la eliminación de un producto
func DeleteProduct(c echo.Context) error {
	// Obtiene el ID del producto de los parámetros de la URL
	id := c.Param("id")

	// Intenta eliminar el producto de la base de datos
	result := db.DB.Delete(&models.Product{}, id)
	if result.Error != nil {
		// Si hay un error al eliminar, devuelve un error 500
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al eliminar el producto"})
	}
	if result.RowsAffected == 0 {
		// Si no se afectó ninguna fila, significa que el producto no existía
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Producto no encontrado"})
	}

	// Devuelve un mensaje de éxito
	return c.JSON(http.StatusOK, map[string]string{"message": "Producto eliminado exitosamente"})
}

// ListProducts maneja la obtención de todos los productos
func ListProducts(c echo.Context) error {
	var products []models.Product

	// Intenta obtener todos los productos de la base de datos
	if err := db.DB.Find(&products).Error; err != nil {
		// Si hay un error al obtener los productos, devuelve un error 500
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener los productos"})
	}

	// Devuelve la lista de productos
	return c.JSON(http.StatusOK, products)
}
