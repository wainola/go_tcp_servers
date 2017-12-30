// Servidor web.
package main

import (
	"net/http"

	c "controllers/controllers"
)

func main() {
	// mux y manejador de recursos estaticos.
	mux := http.NewServeMux()
	// pendiente el servir los archivos estaticos

	// manejo de rutas y funciones
	mux.HandleFunc("/", c.index)

	// server => address to value con operador &
	server := &http.Server{
		Addr:    "http://localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
