package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

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

	mux := http.NewServeMux()

	// Archivos estáticos
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	mux.Handle("/imagenes/", http.StripPrefix("/imagenes/", http.FileServer(http.Dir("./imagenes"))))

	mux.HandleFunc("/", materiaHandler.RenderIndex)
	mux.HandleFunc("/jugar", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			preguntaHandler.RenderJugar(w, r)
		case http.MethodPost:
			preguntaHandler.ProcessJugar(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/crear-pregunta", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			preguntaHandler.RenderForm(w, r)
		case http.MethodPost:
			preguntaHandler.ProcessForm(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/preguntas", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}
		preguntaHandler.RenderPreguntas(w, r)
	})
	mux.HandleFunc("/preguntas/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || !strings.HasSuffix(r.URL.Path, "/eliminar") {
			http.Error(w, "Ruta no válida", http.StatusMethodNotAllowed)
			return
		}
		preguntaHandler.DeleteFromForm(w, r)
	})
	mux.HandleFunc("/editar-pregunta", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			preguntaHandler.RenderEditForm(w, r)
		case http.MethodPost:
			preguntaHandler.ProcessEditForm(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
