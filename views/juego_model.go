package views

import db "ejemplo.com/tp-especial/db/sqlc"

// JuegoState contiene únicamente el estado necesario para reconstruir la vista.
type JuegoState struct {
	Materias    []db.Materium
	Disponibles []db.Materium
	Pregunta    db.Preguntum
	Opciones    []db.Opcion
	Respuesta   string
	Filtro      string
	Error       string
}
