package middleware

import (
	"errors"   // Paquete para crear errores personalizados
	"fmt"      // Paquete para formatear strings
	"net/http" // Paquete para realizar peticiones HTTP y manejar respuestas
	"os"       // Paquete para interactuar con el sistema operativo, aquí usado para variables de entorno

	"github.com/labstack/echo/v4" // Framework web Echo para Go
)

// AuthMiddleware es un middleware que valida el token de autorización
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Obtiene el token del header de autorización
		token := c.Request().Header.Get("Authorization")

		// Verifica si el token está presente
		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token no proporcionado"})
		}

		// Valida el token con el servicio de usuarios
		isValid, err := validateTokenWithUserService(token)
		if err != nil {
			// Si hay un error en la validación, devuelve un error interno del servidor
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al validar el token"})
		}

		// Si el token no es válido, devuelve un error de no autorizado
		if !isValid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token inválido"})
		}

		// Si el token es válido, continúa con el siguiente handler
		return next(c)
	}
}

// validateTokenWithUserService valida el token con el servicio de usuarios
func validateTokenWithUserService(token string) (bool, error) {
	// Construye la URL del servicio de usuarios
	url, err := MakeUserServiceUrl(os.Getenv("APP_USER_SERVICE_URL"))
	if err != nil {
		return false, err
	}

	// Crea una nueva petición GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	// Añade el token al header de autorización
	req.Header.Set("Authorization", token)

	// Crea un cliente HTTP y realiza la petición
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close() // Asegura que el cuerpo de la respuesta se cierre al finalizar

	// Si el código de estado no es OK, considera el token como inválido
	if resp.StatusCode != http.StatusOK {
		return false, nil
	}

	// Si llegamos aquí, el token es válido
	return true, nil
}

// MakeUserServiceUrl crea una URL válida para contactar con el servicio de usuarios
func MakeUserServiceUrl(userServiceHost string) (string, error) {
	var url string

	// Verifica si se ha configurado el host del servicio de usuarios
	if len(userServiceHost) == 0 {
		return url, errors.New("no se ha configurado el host del servicio de usuarios")
	}

	// Construye la URL completa
	url = fmt.Sprintf("%s/validate", userServiceHost)

	return url, nil
}
