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
	materiaHandler := handlers.NewMateriaHandler(queries)
	preguntasHandler := handlers.NewPreguntaHandler(queries)
	// opcionesHandler := handlers.NewOpcionHandler(queries)

	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	// mux.Handle("/imagenes/", http.StripPrefix("/imagenes/", http.FileServer(http.Dir("./imagenes"))))

	mux.Handle("/materias", materiaHandler)
	mux.Handle("/materias/", materiaHandler)
	mux.Handle("/preguntas", preguntasHandler)
	mux.Handle("/preguntas/", preguntasHandler)
	// mux.Handle("/opciones", opcionesHandler)
	// mux.Handle("/opciones/", opcionesHandler)

	log.Println("Servidor escuchando en http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
