package middleware_test

import (
	"fmt"
	"product-service/middleware"
	"testing"
)

func TestMakeUserServiceUrl(t *testing.T) {

	_, err := middleware.MakeUserServiceUrl("")
	if err == nil {
		t.Errorf("Se esperaba un error cuando se llama a MakeUserServiceUrl con un host vacío")
	}

	urlToTest := "http://localhost:8080"

	url, err := middleware.MakeUserServiceUrl(urlToTest)
	if err != nil {
		t.Errorf("Se esperaba un error cuando se llama a MakeUserServiceUrl con un host válido")
	}

	if fmt.Sprintf("%s/validate", urlToTest) != url {
		t.Errorf("El resultado de MakeUserServiceUrl debería ser igual al host proporcionado")
	}

}
