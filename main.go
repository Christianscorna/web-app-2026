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

	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	mux.Handle("/imagenes/", http.StripPrefix("/imagenes/", http.FileServer(http.Dir("./imagenes"))))
	mux.Handle("/materias", materiashandler)
	mux.Handle("/materias/", materiashandler)
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Servidor escuchando en http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}
