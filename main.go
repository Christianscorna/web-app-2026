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

  	// la coneccion se cierra
	defer conn.Close()

	queries := db.New(conn)
	materiashandler := handlers.NewMateriaHandler(queries)

	fileServer := http.FileServer(http.Dir("./static")) // luego hay que ver como mapear los archivos html con los handlers

	http.Handle("/", fileServer)
	
	http.Handle("/materias", materiashandler)
	http.Handle("/materias/", materiashandler)
	http.Handle("/materias/{id}", materiashandler)

	log.Println("Servidor escuchando en http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
