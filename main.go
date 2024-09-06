package main

import (
	"product-service/config"
	"product-service/db"
	"product-service/handlers"
	"product-service/middleware"
	"time"

	"github.com/labstack/echo/v4"
	echo_middlewares "github.com/labstack/echo/v4/middleware"
)

func main() {
	conf := config.Config{}
	conf.LoadConfig()

	if conf.APPEnv != "development" {
		time.Sleep(time.Second * 5)
	}
	db.ConnectDatabase(&conf)

	e := echo.New()

	e.Use(echo_middlewares.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Product service is running! v1.0.1")
	})

	// Grupo de rutas protegidas
	protected := e.Group("")

	protected.Use(middleware.AuthMiddleware)

	// Rutas protegidas
	protected.POST("/products", handlers.CreateProduct)
	protected.GET("/products/:id", handlers.GetProduct)
	protected.PUT("/products/:id", handlers.UpdateProduct)
	protected.DELETE("/products/:id", handlers.DeleteProduct)
	protected.GET("/products", handlers.ListProducts)

	e.Logger.Fatal(e.Start(":8081"))

}
