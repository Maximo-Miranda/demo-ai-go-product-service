package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token no proporcionado"})
		}

		// Aquí llamaremos al servicio de usuarios para validar el token
		isValid, err := validateTokenWithUserService(token)
		if err != nil {
			panic(err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al validar el token"})
		}

		if !isValid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token inválido"})
		}

		return next(c)
	}
}

func validateTokenWithUserService(token string) (bool, error) {
	userServiceURL := os.Getenv("APP_USER_SERVICE_URL")

	fmt.Println("USER SERVICE URL", userServiceURL)

	url := fmt.Sprintf("%s/validate", userServiceURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, nil
	}

	return true, nil
}
