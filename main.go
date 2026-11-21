package main

import (
	"fmt"
	"log"
	"net/http"

	. "ejemplo.com/tp-especial/db"
	"ejemplo.com/tp-especial/db/sqlc"
	"ejemplo.com/tp-especial/handlers"
)

func main() {
	conn, err := ConnectDB()
	if err != nil {
		fmt.Printf("Error al conectar a la base de datos: %s\n", err)
		return
	}
	defer conn.Close()

	queries := db.New(conn)

	// Instanciamos los manejadores
	materiaHandler := handlers.NewMateriaHandler(queries)
	preguntaHandler := handlers.NewPreguntaHandler(queries)
	router := handlers.NewRouter(materiaHandler, preguntaHandler)

	log.Println("Servidor escuchando en http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
