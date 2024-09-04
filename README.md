# Documentación del Servicio de Productos (@product-service)

## Descripción General

Este proyecto es un microservicio de gestión de productos desarrollado en Go, utilizando el framework Echo para la creación de APIs RESTful. El servicio proporciona funcionalidades para la gestión de productos, incluyendo creación, actualización, eliminación y consulta de productos.

## Estructura del Proyecto

El proyecto sigue una estructura modular típica de aplicaciones Go:

- `config/`: Contiene la configuración de la aplicación.
- `db/`: Maneja la conexión y operaciones con la base de datos.
- `handlers/`: Contiene los manejadores de las rutas HTTP.
- `middleware/`: Incluye middlewares personalizados.
- `models/`: Define las estructuras de datos utilizadas en la aplicación.
- `main.go`: Punto de entrada de la aplicación.

## Configuración

La configuración se maneja a través de variables de entorno y un archivo `.env` para entornos de desarrollo. Las principales configuraciones incluyen:

- `DB_CONNECTION_STRING`: Cadena de conexión a la base de datos PostgreSQL.
- `APP_ENV`: Entorno de la aplicación (development, production, etc.).
- `APP_USER_SERVICE_URL`: URL del servicio de usuarios.

## Rutas Principales

1. Crear Producto:
   - Ruta: `POST /products`
   - Funcionalidad: Permite crear un nuevo producto en el sistema.

2. Obtener Producto:
   - Ruta: `GET /products/{id}`
   - Funcionalidad: Obtiene la información de un producto específico.

3. Listar Productos:
   - Ruta: `GET /products`
   - Funcionalidad: Obtiene una lista de todos los productos registrados.

4. Actualizar Producto:
   - Ruta: `PUT /products/{id}`
   - Funcionalidad: Actualiza la información de un producto existente.

5. Eliminar Producto:
   - Ruta: `DELETE /products/{id}`
   - Funcionalidad: Elimina un producto del sistema.

## Modelos de Datos

El modelo principal es `Product`, que incluye campos como:

```5:11:user-product-service/models/product.go
type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}
```
## Seguridad

- Se utiliza un middleware de autenticación para proteger rutas específicas el cual es procesado por el servicio de usuarios.

## Despliegue

El proyecto incluye configuraciones para despliegue utilizando Docker y Docker Compose. Los archivos relevantes son:

- `Dockerfile`: Para la construcción de la imagen Docker del servicio.
- `docker-compose.yml`: Para orquestar el servicio junto con la base de datos PostgreSQL.

## Desarrollo y Pruebas

Para el desarrollo local:

1. Clonar el repositorio.
2. Copiar el archivo `.env.example` a `.env` y configurar las variables de entorno.
3. Ejecutar `go mod download` para instalar las dependencias.
4. Usar `go run main.go` para iniciar el servidor de desarrollo.

Para pruebas, se incluye un flujo de CI/CD en GitHub Actions que ejecuta pruebas automáticas en cada pull request.

## Notas Adicionales

Este proyecto es para fines educativos y demuestra conceptos como:
- Desarrollo de microservicios en Go
- Uso de frameworks web como Echo
- Implementación de autenticación y autorización
- Manejo de bases de datos con GORM
- Configuración de CI/CD con GitHub Actions

Se recomienda revisar y mejorar las prácticas de seguridad antes de utilizar en un entorno de producción real.


