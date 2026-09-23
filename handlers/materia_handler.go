package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"ejemplo.com/tp-especial/db/sqlc"
)

type MateriaHandler struct {
	queries *db.Queries
}

func NewMateriaHandler(queries *db.Queries) *MateriaHandler {
	return &MateriaHandler{queries: queries}
}

// Auxiliar para modularizar el envío de errores o set de estado
func throwingError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// Auxiliar para extraer el ID del path
func extractIDFromPath(path string) (int, bool, error) {
	if path == "/materias" || path == "/materias/" {
		return 0, false, nil
	}

	// TrimPrefix hace que si el path es "/materias/5" se quede con "5"
	idText := strings.TrimPrefix(path, "/materias/")
	fmt.Println("ID extraído del path:", idText)

	// Atoi convierte el string a int, si no puede devuelve un error
	id, err := strconv.Atoi(idText)
	if err != nil || id <= 0 {
		return 0, true, fmt.Errorf("ID inválido")
	}
	
	return id, true, nil
}

func (h *MateriaHandler) getMaterias(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	materias, err := h.queries.ListMaterias(r.Context())
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al obtener las materias")
		return
	}

	err = json.NewEncoder(w).Encode(materias)
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al codificar las materias")
		return
	}
}

func (h *MateriaHandler) getMateriaByID(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")
	materia, err := h.queries.GetMateria(r.Context(), int32(id))
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al obtener la materia")
		return
	}

	err = json.NewEncoder(w).Encode(materia)
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al codificar la materia")
		return
	}
}

func (h *MateriaHandler) createMateria(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var materia db.CreateMateriaParams
	err := json.NewDecoder(r.Body).Decode(&materia)
	if err != nil {
		throwingError(w, http.StatusBadRequest, "Error al decodificar la materia")
		return
	}

	createdMateria, err := h.queries.CreateMateria(r.Context(), materia)
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al crear la materia")
		return
	}

	err = json.NewEncoder(w).Encode(createdMateria)
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al codificar la materia creada")
		return
	}
}

func (h *MateriaHandler) updateMateria(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")
	var materia db.UpdateMateriaParams
	err := json.NewDecoder(r.Body).Decode(&materia)
	if err != nil {
		throwingError(w, http.StatusBadRequest, "Error al decodificar la materia")
		return
	}

	materia.IDMateria = int32(id)
	err = h.queries.UpdateMateria(r.Context(), materia)
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al actualizar la materia")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *MateriaHandler) deleteMateria(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")
	err := h.queries.DeleteMateria(r.Context(), int32(id))
	if err != nil {
		throwingError(w, http.StatusInternalServerError, "Error al eliminar la materia")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *MateriaHandler) handleMaterias(w http.ResponseWriter, r *http.Request) {
	id, hasID, err := extractIDFromPath(r.URL.Path)
	if err != nil {
		throwingError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	switch r.Method {
	case http.MethodGet:
		if !hasID {
			h.getMaterias(w, r)
		} else {
			h.getMateriaByID(w, r, id)
		}

	case http.MethodPost:
		h.createMateria(w, r)

	case http.MethodPut:
		if !hasID {
			throwingError(w, http.StatusBadRequest, "ID requerido para actualizar")
			return
		}
		h.updateMateria(w, r, id)

	case http.MethodDelete:
		if !hasID {
			throwingError(w, http.StatusBadRequest, "ID requerido para eliminar")
			return
		}
		h.deleteMateria(w, r, id)

	default:
		throwingError(w, http.StatusMethodNotAllowed, "Método no permitido")
	}
}

func (h *MateriaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handleMaterias(w, r)
}
