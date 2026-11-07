package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"ejemplo.com/tp-especial/db/sqlc"
	"ejemplo.com/tp-especial/views"
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

// En handlers/pregunta_handler.go

// Muestra la vista del formulario
func (h *PreguntaHandler) RenderForm(w http.ResponseWriter, r *http.Request) {
	materiaPrevia := r.URL.Query().Get("materia")
	err := views.CrearPreguntaPage(materiaPrevia).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error al renderizar el formulario", http.StatusInternalServerError)
	}
}

// Procesa los datos enviados por el <form>
func (h *PreguntaHandler) ProcessForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	materiaNombre := r.FormValue("materia")
	enunciado := r.FormValue("enunciado")
	opcionCorrectaNum := r.FormValue("opcion_correcta")

	// 1. Buscar o crear materia por nombre
	materias, err := h.queries.ListMaterias(r.Context())
	if err != nil {
		http.Error(w, "Error al obtener las materias", http.StatusInternalServerError)
		return
	}

	var materia db.Materium
	encontrada := false

	for _, m := range materias {
		if m.Nombre == materiaNombre {
			materia = m
			encontrada = true
			break
		}
	}

	if !encontrada {
		materia, err = h.queries.CreateMateria(r.Context(), db.CreateMateriaParams{
			Nombre: materiaNombre,
		})
		if err != nil {
			http.Error(w, "Error al guardar la materia", http.StatusInternalServerError)
			return
		}
	}

	// 2. Crear la pregunta
	pregunta, err := h.queries.CreatePregunta(r.Context(), db.CreatePreguntaParams{
		IDMateria: materia.IDMateria,
		Enunciado: enunciado,
	})
	if err != nil {
		http.Error(w, "Error al guardar la pregunta", http.StatusInternalServerError)
		return
	}

	// 3. Guardar las 4 opciones
	opciones := []string{
		r.FormValue("opcion_1"),
		r.FormValue("opcion_2"),
		r.FormValue("opcion_3"),
		r.FormValue("opcion_4"),
	}

	for i, texto := range opciones {
		esCorrecta := strconv.Itoa(i+1) == opcionCorrectaNum
		_, err := h.queries.CreateOpcion(r.Context(), db.CreateOpcionParams{
			IDPregunta: pregunta.IDPregunta,
			Texto:      texto,
			EsCorrecta: esCorrecta,
		})
		if err != nil {
			http.Error(w, "Error al guardar las opciones", http.StatusInternalServerError)
			return
		}
	}

	// 4. Redirección PRG a la pantalla principal
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// RenderPreguntas muestra la pantalla SSR para administrar las preguntas de una materia.
func (h *PreguntaHandler) RenderPreguntas(w http.ResponseWriter, r *http.Request) {
	materias, err := h.queries.ListMaterias(r.Context())
	if err != nil {
		http.Error(w, "Error al obtener las materias", http.StatusInternalServerError)
		return
	}

	nombreMateria := strings.TrimSpace(r.URL.Query().Get("materia"))
	var materia db.Materium
	for _, candidata := range materias {
		if strings.EqualFold(candidata.Nombre, nombreMateria) ||
			strconv.FormatInt(int64(candidata.IDMateria), 10) == nombreMateria {
			materia = candidata
			break
		}
	}

	var preguntas []db.Preguntum
	if materia.IDMateria != 0 {
		preguntas, err = h.queries.ListPreguntasByMateria(r.Context(), materia.IDMateria)
		if err != nil {
			http.Error(w, "Error al obtener las preguntas", http.StatusInternalServerError)
			return
		}
	}

	if err := views.GestionPreguntasPage(materias, materia, preguntas).Render(r.Context(), w); err != nil {
		http.Error(w, "Error al renderizar la gestión de preguntas", http.StatusInternalServerError)
	}
}

func (h *PreguntaHandler) DeleteFromForm(w http.ResponseWriter, r *http.Request) {
	id, hasID, err := extractPreguntaIDFromPath(strings.TrimSuffix(r.URL.Path, "/eliminar"))
	if err != nil || !hasID || id <= 0 {
		http.Error(w, "ID de pregunta inválido", http.StatusBadRequest)
		return
	}

	if err := h.queries.DeletePregunta(r.Context(), int32(id)); err != nil {
		http.Error(w, "Error al eliminar la pregunta", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/preguntas", http.StatusSeeOther)
}

func (h *PreguntaHandler) RenderEditForm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		http.Error(w, "ID de pregunta inválido", http.StatusBadRequest)
		return
	}

	pregunta, err := h.queries.GetPregunta(r.Context(), int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "Error al obtener la pregunta", http.StatusInternalServerError)
		return
	}
	opciones, err := h.queries.ListOpcionesByPregunta(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Error al obtener las opciones", http.StatusInternalServerError)
		return
	}
	materias, err := h.queries.ListMaterias(r.Context())
	if err != nil {
		http.Error(w, "Error al obtener las materias", http.StatusInternalServerError)
		return
	}

	if err := views.EditarPreguntaPage(pregunta, opciones, materias).Render(r.Context(), w); err != nil {
		http.Error(w, "Error al renderizar la edición", http.StatusInternalServerError)
	}
}

func (h *PreguntaHandler) ProcessEditForm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		http.Error(w, "ID de pregunta inválido", http.StatusBadRequest)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}
	idMateria, err := strconv.Atoi(r.FormValue("id_materia"))
	if err != nil || idMateria <= 0 {
		http.Error(w, "Materia inválida", http.StatusBadRequest)
		return
	}
	if err := h.queries.UpdatePregunta(r.Context(), db.UpdatePreguntaParams{
		IDPregunta: int32(id),
		Enunciado:  r.FormValue("enunciado"),
		IDMateria:  int32(idMateria),
	}); err != nil {
		http.Error(w, "Error al actualizar la pregunta", http.StatusInternalServerError)
		return
	}

	for i := 1; i <= 4; i++ {
		opcionID, err := strconv.Atoi(r.FormValue(fmt.Sprintf("opcion_id_%d", i)))
		if err != nil || opcionID <= 0 {
			continue
		}
		err = h.queries.UpdateOpcion(r.Context(), db.UpdateOpcionParams{
			IDOpcion:   int32(opcionID),
			Texto:      r.FormValue(fmt.Sprintf("opcion_%d", i)),
			EsCorrecta: r.FormValue("opcion_correcta") == strconv.Itoa(i),
		})
		if err != nil {
			http.Error(w, "Error al actualizar las opciones", http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/preguntas", http.StatusSeeOther)
}
