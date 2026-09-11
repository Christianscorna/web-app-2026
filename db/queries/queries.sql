-- db/queries/queries.sql

-- === Etidad MATERIA ===
-- name: CreateMateria :one
INSERT INTO materia (nombre, descripcion)
VALUES ($1, $2)
RETURNING id_materia, nombre, descripcion;

-- name: GetMateria :one
SELECT id_materia, nombre, descripcion
FROM materia
WHERE id_materia = $1;

-- name: ListMaterias :many
SELECT id_materia, nombre, descripcion
FROM materia
ORDER BY nombre;

-- name: UpdateMateria :exec
UPDATE materia
SET nombre = $2, descripcion = $3
WHERE id_materia = $1;

-- name: DeleteMateria :exec
DELETE FROM materia
WHERE id_materia = $1;

-- === Entidad PREGUNTA ===
-- name: CreatePregunta :one
INSERT INTO pregunta (enunciado, id_materia)
VALUES ($1, $2)
RETURNING id_pregunta, enunciado, id_materia;

-- name: GetPregunta :one
SELECT id_pregunta, enunciado, id_materia
FROM pregunta
WHERE id_pregunta = $1;

-- name: ListPreguntasByMateria :many
SELECT id_pregunta, enunciado, id_materia
FROM pregunta
WHERE id_materia = $1
ORDER BY id_pregunta;

-- name: UpdatePregunta :exec
UPDATE pregunta
SET enunciado = $2, id_materia = $3
WHERE id_pregunta = $1;

-- name: DeletePregunta :exec
DELETE FROM pregunta
WHERE id_pregunta = $1;

-- === Entidad OPCION ===
-- name: CreateOpcion :one
INSERT INTO opcion (texto, es_correcta, id_pregunta)
VALUES ($1, $2, $3)
RETURNING id_opcion, texto, es_correcta, id_pregunta;

-- name: GetOpcion :one
SELECT id_opcion, texto, es_correcta, id_pregunta
FROM opcion
WHERE id_opcion = $1;

-- name: ListOpcionesByPregunta :many
SELECT id_opcion, texto, es_correcta, id_pregunta
FROM opcion
WHERE id_pregunta = $1
ORDER BY id_opcion;

-- name: UpdateOpcion :exec
UPDATE opcion
SET texto = $2, es_correcta = $3
WHERE id_opcion = $1;

-- name: DeleteOpcion :exec
DELETE FROM opcion
WHERE id_opcion = $1;
