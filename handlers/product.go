package handlers

import (
	"net/http"
	"product-service/db"
	"product-service/models"

	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos de solicitud inválidos"})
	}

	db.DB.Create(&product)
	return c.JSON(http.StatusCreated, product)
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Producto no encontrado"})
	}
	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Producto no encontrado"})
	}

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos de solicitud inválidos"})
	}

	db.DB.Save(&product)
	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Producto no encontrado"})
	}

	db.DB.Delete(&product)
	return c.JSON(http.StatusOK, map[string]string{"message": "Producto eliminado exitosamente"})
}

func ListProducts(c echo.Context) error {
	var products []models.Product
	db.DB.Find(&products)
	return c.JSON(http.StatusOK, products)
}
