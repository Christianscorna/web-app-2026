package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"ejemplo.com/tp-especial/db/sqlc"
)

type PreguntaHandler struct {
	queries *db.Queries
}

func NewPreguntaHandler(queries *db.Queries) *PreguntaHandler {
	return &PreguntaHandler{queries: queries}
}

func (h *PreguntaHandler) getPreguntas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idMateriaText := r.URL.Query().Get("id_materia")
	if idMateriaText == "" {
		throwingError(w, http.StatusBadRequest, "Parámetro id_materia requerido")
		return
	}

	idMateria, err := strconv.Atoi(idMateriaText)
	if err != nil || idMateria <= 0 {
		throwingError(w, http.StatusBadRequest, "ID de materia inválido")
		return
	}

	preguntas, err := h.queries.ListPreguntasByMateria(r.Context(), int32(idMateria))
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al obtener las preguntas")
		return
	}

	if err := json.NewEncoder(w).Encode(preguntas); err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al codificar las preguntas")
	}
}

func (h *PreguntaHandler) getPreguntaByID(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")
	pregunta, err := h.queries.GetPregunta(r.Context(), int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			throwingError(w, http.StatusNotFound, "Pregunta no encontrada")
			return
		}
		throwingError(w, http.StatusInternalServerError, "Error al obtener la pregunta")
		return
	}

	if err := json.NewEncoder(w).Encode(pregunta); err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al codificar la pregunta")
	}
}

func (h *PreguntaHandler) createPregunta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pregunta db.CreatePreguntaParams
	if err := json.NewDecoder(r.Body).Decode(&pregunta); err != nil {
		throwingError(w, http.StatusBadRequest, "Error al decodificar la pregunta")
		return
	}

	createdPregunta, err := h.queries.CreatePregunta(r.Context(), pregunta)
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al crear la pregunta")
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdPregunta); err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al codificar la pregunta creada")
	}
}

func (h *PreguntaHandler) updatePregunta(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")
	var pregunta db.UpdatePreguntaParams
	if err := json.NewDecoder(r.Body).Decode(&pregunta); err != nil {
		throwingError(w, http.StatusBadRequest, "Error al decodificar la pregunta")
		return
	}

	pregunta.IDPregunta = int32(id)
	if err := h.queries.UpdatePregunta(r.Context(), pregunta); err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al actualizar la pregunta")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *PreguntaHandler) deletePregunta(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")
	if err := h.queries.DeletePregunta(r.Context(), int32(id)); err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al eliminar la pregunta")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Este helper limpia la ruta /preguntas y verifica si existe un ID numérico a continuación.
func extractPreguntaIDFromPath(path string) (int, bool, error) {
	trimmed := strings.TrimPrefix(path, "/preguntas")
	trimmed = strings.Trim(trimmed, "/")

	if trimmed == "" {
		return 0, false, nil
	}

	id, err := strconv.Atoi(trimmed)
	if err != nil {
		return 0, true, err
	}

	return id, true, nil
}

func (h *PreguntaHandler) handlePreguntas(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			id, hasID, err := extractPreguntaIDFromPath(r.URL.Path)
			if err != nil {
				throwingError(w, http.StatusBadRequest, "ID de pregunta inválido")
				return
			}

			if hasID {
				h.getPreguntaByID(w, r, id)
			} else {
				h.getPreguntas(w, r)
			}

		case http.MethodPost:
			if r.URL.Path == "/preguntas" || r.URL.Path == "/preguntas/" {
				h.createPregunta(w, r)
				return
			}
			throwingError(w, http.StatusBadRequest, "Ruta no válida para creación de pregunta")

		case http.MethodPut:
			id, hasID, err := extractPreguntaIDFromPath(r.URL.Path)
			if err != nil {
				throwingError(w, http.StatusBadRequest, "ID de pregunta inválido")
				return
			}
			if !hasID {
				throwingError(w, http.StatusBadRequest, "ID requerido para actualizar")
				return
			}
			h.updatePregunta(w, r, id)

		case http.MethodDelete:
			id, hasID, err := extractPreguntaIDFromPath(r.URL.Path)
			if err != nil {
				throwingError(w, http.StatusBadRequest, "ID de pregunta inválido")
				return
			}
			if !hasID {
				throwingError(w, http.StatusBadRequest, "ID requerido para eliminar")
				return
			}
			h.deletePregunta(w, r, id)

		default:
			throwingError(w, http.StatusMethodNotAllowed, "Método no permitido")
	}
}

func (h *PreguntaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handlePreguntas(w, r)
}
