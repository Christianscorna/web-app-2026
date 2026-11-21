package handlers

import (
	"net/http"
	"strings"
)

func NewRouter(materiaHandler *MateriaHandler, preguntaHandler *PreguntaHandler) http.Handler {
	mux := http.NewServeMux()

	mux.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))),
	)

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
		switch r.Method {
		case http.MethodDelete:
			preguntaHandler.DeletePreguntaFromURL(w, r)
		case http.MethodPost:
			if !strings.HasSuffix(r.URL.Path, "/eliminar") {
				http.Error(w, "Ruta no válida", http.StatusMethodNotAllowed)
				return
			}
			preguntaHandler.DeleteFromForm(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
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

	return mux
}
