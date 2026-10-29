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

type OpcionHandler struct {
	queries *db.Queries
}

func NewOpcionHandler(queries *db.Queries) *OpcionHandler {
	return &OpcionHandler{queries: queries}
}

func (h *OpcionHandler) getOpciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idPregunta, err := strconv.Atoi(r.URL.Query().Get("id_pregunta"))
	if err != nil || idPregunta <= 0 {
		throwingError(w, http.StatusBadRequest, "Parámetro id_pregunta requerido y debe ser válido")
		return
	}

	opciones, err := h.queries.ListOpcionesByPregunta(r.Context(), int32(idPregunta))
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al obtener las opciones")
		return
	}

	err = json.NewEncoder(w).Encode(opciones)
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al codificar las opciones")
		return
	}
}

func (h *OpcionHandler) getOpcionByID(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")
	opcion, err := h.queries.GetOpcion(r.Context(), int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			throwingError(w, http.StatusNotFound, "Opcion no encontrada")
			return
		}
		throwingError(w, http.StatusInternalServerError, "Error al obtener la opcion")
		return
	}

	err = json.NewEncoder(w).Encode(opcion)
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al codificar la opcion")
		return
	}
}

func (h *OpcionHandler) createOpcion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var opcion db.CreateOpcionParams
	err := json.NewDecoder(r.Body).Decode(&opcion)
	if err != nil {
		throwingError(w, http.StatusBadRequest, "Error al decodificar la opcion")
		return
	}

	createdOpcion, err := h.queries.CreateOpcion(r.Context(), opcion)
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al crear la opcion")
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdOpcion)
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al codificar la opcion creada")
		return
	}
}

func (h *OpcionHandler) updateOpcion(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")
	var opcion db.UpdateOpcionParams
	err := json.NewDecoder(r.Body).Decode(&opcion)
	if err != nil {
		throwingError(w, http.StatusBadRequest, "Error al decodificar la opcion")
		return
	}

	opcion.IDOpcion = int32(id)
	err = h.queries.UpdateOpcion(r.Context(), opcion)
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al actualizar la opcion")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *OpcionHandler) deleteOpcion(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")
	err := h.queries.DeleteOpcion(r.Context(), int32(id))
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al eliminar la opcion")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *OpcionHandler) handleOpciones(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/opciones" && r.URL.Path != "/opciones/" && !strings.HasPrefix(r.URL.Path, "/opciones/") {
		throwingError(w, http.StatusNotFound, "Ruta no encontrada")
		return
	}

	idText := strings.Trim(strings.TrimPrefix(r.URL.Path, "/opciones"), "/")
	id, hasID := 0, idText != ""
	if hasID {
		var err error
		id, err = strconv.Atoi(idText)
		if err != nil || id <= 0 {
			throwingError(w, http.StatusBadRequest, "ID inválido")
			return
		}
	}

	switch r.Method {
	case http.MethodGet:
		if !hasID {
			h.getOpciones(w, r)
		} else {
			h.getOpcionByID(w, r, id)
		}

	case http.MethodPost:
		if hasID {
			throwingError(w, http.StatusMethodNotAllowed, "POST solo está disponible en /opciones")
			return
		}
		h.createOpcion(w, r)

	case http.MethodPut:
		if !hasID {
			throwingError(w, http.StatusBadRequest, "ID requerido para actualizar")
			return
		}
		h.updateOpcion(w, r, id)

	case http.MethodDelete:
		if !hasID {
			throwingError(w, http.StatusBadRequest, "ID requerido para eliminar")
			return
		}
		h.deleteOpcion(w, r, id)

	default:
		throwingError(w, http.StatusMethodNotAllowed, "Método no permitido")
	}
}

func (h *OpcionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handleOpciones(w, r)
}
