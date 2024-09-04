// Package middleware_test contiene las pruebas para el paquete middleware
package middleware_test

import (
	"fmt"                        // Paquete para formatear strings
	"product-service/middleware" // Importa el paquete middleware local que se está probando
	"testing"                    // Paquete para escribir y ejecutar pruebas unitarias
)

// TestMakeUserServiceUrl prueba la función MakeUserServiceUrl
func TestMakeUserServiceUrl(t *testing.T) {

	// Prueba con un host vacío
	_, err := middleware.MakeUserServiceUrl("")
	if err == nil {
		// Si no hay error cuando se espera uno, la prueba falla
		t.Errorf("Se esperaba un error cuando se llama a MakeUserServiceUrl con un host vacío")
	}

	// URL de prueba
	urlToTest := "http://localhost:8080"

	// Prueba con un host válido
	url, err := middleware.MakeUserServiceUrl(urlToTest)
	if err != nil {
		// Si hay un error cuando no se espera, la prueba falla
		t.Errorf("No se esperaba un error cuando se llama a MakeUserServiceUrl con un host válido")
	}

	// Verifica que la URL resultante sea la esperada
	expectedUrl := fmt.Sprintf("%s/validate", urlToTest)
	if expectedUrl != url {
		// Si la URL resultante no es la esperada, la prueba falla
		t.Errorf("El resultado de MakeUserServiceUrl debería ser %s, pero se obtuvo %s", expectedUrl, url)
	}

}
