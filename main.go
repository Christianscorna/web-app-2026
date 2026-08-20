package main

import (
  "fmt"
  "net/http"
)

func main() {

  // 1. Crea un (handler) que sirve archivos estáticos del dir static.
  fileServer := http.FileServer(http.Dir("./static"))

	// 2. Manejador de ruta raíz "/"
  http.Handle("/", fileServer)

	// 3. Usamos el puerto que pide el enunciado
	port := ":8080"
  fmt.Printf("Servidor ESTÁTICO escuchando en http://localhost%s\n", port)

  // 4. Si no hay errores, inicia servidor. 
	err := http.ListenAndServe(port, nil)
  if err != nil {
    fmt.Printf("Error al iniciar el servidor: %s\n", err)
  }
}
