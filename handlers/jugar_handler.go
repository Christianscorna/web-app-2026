package handlers

import (
	"crypto/rand"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	db "ejemplo.com/tp-especial/db/sqlc"
	"ejemplo.com/tp-especial/views"
)

const maxMateriasJuego = 5

func (h *PreguntaHandler) RenderJugar(w http.ResponseWriter, r *http.Request) {
	materias, err := h.queries.ListMaterias(r.Context())
	if err != nil {
		http.Error(w, "Error al obtener las materias", http.StatusInternalServerError)
		return
	}
	estado := views.JuegoState{
		Materias: juegoStateFromIDs(materias, r.URL.Query()["materias"]),
	}
	estado.Filtro = strings.TrimSpace(r.URL.Query().Get("q"))
	estado.Disponibles = filtrarMaterias(materias, estado.Filtro)
	if err := views.VistaDeJuego(estado).Render(r.Context(), w); err != nil {
		http.Error(w, "Error al renderizar el juego", http.StatusInternalServerError)
	}
}

func (h *PreguntaHandler) ProcessJugar(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar el juego", http.StatusBadRequest)
		return
	}
	materias, err := h.queries.ListMaterias(r.Context())
	if err != nil {
		http.Error(w, "Error al obtener las materias", http.StatusInternalServerError)
		return
	}
	seleccionadas := juegoStateFromIDs(materias, r.Form["materias"])
	estado := views.JuegoState{Materias: seleccionadas}
	estado.Disponibles = filtrarMaterias(materias, strings.TrimSpace(r.FormValue("q")))
	estado.Filtro = strings.TrimSpace(r.FormValue("q"))

	accion := r.FormValue("accion")
	if accion == "" {
		accion = "girar"
	}
	switch accion {
	case "agregar":
		estado.Materias = agregarMateria(seleccionadas, r.FormValue("id_materia"), materias)
	case "quitar":
		estado.Materias = quitarMateria(seleccionadas, r.FormValue("id_materia"))
	case "girar":
		estado.Materias = seleccionadas
		h.cargarPregunta(r, &estado)
	case "responder":
		estado.Materias = seleccionadas
		h.procesarRespuesta(r, &estado)
	default:
		estado.Error = "Acción de juego inválida."
	}

	if estado.Error == "" && len(estado.Materias) > 0 && estado.Respuesta == "" {
		estado.Disponibles = filtrarMaterias(materias, estado.Filtro)
	}
	if err := views.VistaDeJuego(estado).Render(r.Context(), w); err != nil {
		http.Error(w, "Error al renderizar el juego", http.StatusInternalServerError)
	}
}

func (h *PreguntaHandler) cargarPregunta(r *http.Request, estado *views.JuegoState) {
	if len(estado.Materias) == 0 {
		estado.Error = "Agregá al menos una materia antes de girar."
		return
	}
	indice, err := rand.Int(rand.Reader, big.NewInt(int64(len(estado.Materias))))
	if err != nil {
		estado.Error = "No se pudo girar la ruleta."
		return
	}
	materia := estado.Materias[indice.Int64()]
	preguntas, err := h.queries.ListPreguntasByMateria(r.Context(), materia.IDMateria)
	if err != nil {
		estado.Error = "No se pudieron obtener las preguntas."
		return
	}
	if len(preguntas) == 0 {
		estado.Error = "La materia elegida todavía no tiene preguntas."
		return
	}
	preguntaIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(preguntas))))
	if err != nil {
		estado.Error = "No se pudo elegir una pregunta."
		return
	}
	estado.Pregunta = preguntas[preguntaIndex.Int64()]
	estado.Opciones, err = h.queries.ListOpcionesByPregunta(r.Context(), estado.Pregunta.IDPregunta)
	if err != nil {
		estado.Error = "No se pudieron obtener las opciones."
	}
}

func (h *PreguntaHandler) procesarRespuesta(r *http.Request, estado *views.JuegoState) {
	idPregunta, err := strconv.Atoi(r.FormValue("id_pregunta"))
	if err != nil || idPregunta <= 0 {
		estado.Error = "Pregunta inválida."
		return
	}
	estado.Pregunta, err = h.queries.GetPregunta(r.Context(), int32(idPregunta))
	if err != nil {
		estado.Error = "No se pudo recuperar la pregunta."
		return
	}
	if !materiaSeleccionada(estado.Pregunta.IDMateria, estado.Materias) {
		estado.Error = "La pregunta no pertenece a una materia seleccionada."
		return
	}
	estado.Opciones, err = h.queries.ListOpcionesByPregunta(r.Context(), estado.Pregunta.IDPregunta)
	if err != nil {
		estado.Error = "No se pudieron obtener las opciones."
		return
	}
	idOpcion, err := strconv.Atoi(r.FormValue("id_opcion"))
	if err != nil {
		estado.Error = "Seleccioná una opción válida."
		return
	}
	for _, opcion := range estado.Opciones {
		if opcion.IDOpcion == int32(idOpcion) {
			if opcion.EsCorrecta {
				estado.Respuesta = "¡Muy bien! Has acertado la respuesta."
			} else {
				estado.Respuesta = "Incorrecto. Probá otra pregunta."
			}
			return
		}
	}
	estado.Error = "La opción no pertenece a la pregunta."
}

func materiaSeleccionada(id int32, materias []db.Materium) bool {
	for _, materia := range materias {
		if materia.IDMateria == id {
			return true
		}
	}
	return false
}

func juegoStateFromIDs(materias []db.Materium, ids []string) []db.Materium {
	seleccionadas := make([]db.Materium, 0, len(ids))
	seen := make(map[int32]bool)
	for _, rawID := range ids {
		id, err := strconv.Atoi(rawID)
		if err != nil || id <= 0 {
			continue
		}
		for _, materia := range materias {
			if materia.IDMateria == int32(id) && !seen[materia.IDMateria] && len(seleccionadas) < maxMateriasJuego {
				seleccionadas = append(seleccionadas, materia)
				seen[materia.IDMateria] = true
				break
			}
		}
	}
	return seleccionadas
}

func agregarMateria(actuales []db.Materium, rawID string, catalogo []db.Materium) []db.Materium {
	id, err := strconv.Atoi(rawID)
	if err != nil || len(actuales) >= maxMateriasJuego {
		return actuales
	}
	for _, materia := range actuales {
		if materia.IDMateria == int32(id) {
			return actuales
		}
	}
	for _, materia := range catalogo {
		if materia.IDMateria == int32(id) {
			return append(actuales, materia)
		}
	}
	return actuales
}

func quitarMateria(actuales []db.Materium, rawID string) []db.Materium {
	id, err := strconv.Atoi(rawID)
	if err != nil {
		return actuales
	}
	resultado := make([]db.Materium, 0, len(actuales))
	for _, materia := range actuales {
		if materia.IDMateria != int32(id) {
			resultado = append(resultado, materia)
		}
	}
	return resultado
}

func filtrarMaterias(materias []db.Materium, filtro string) []db.Materium {
	if filtro == "" {
		return materias
	}
	filtro = strings.ToLower(filtro)
	resultado := make([]db.Materium, 0, len(materias))
	for _, materia := range materias {
		if strings.Contains(strings.ToLower(materia.Nombre), filtro) {
			resultado = append(resultado, materia)
		}
	}
	return resultado
}
