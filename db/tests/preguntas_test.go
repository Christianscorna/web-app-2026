package tests

import (
  "context"
  "database/sql"
  "errors"
  "testing"
  "time"

  _ "github.com/lib/pq"
  // esto es para poder tener las pruebas en un subdir a parte
  . "ejemplo.com/tp-especial/db/sqlc" // dot import para no tener que poner sqlc. por cada Create... Update...
)

func TestPregunta_CRUD(t *testing.T) {
  connStr := "postgres://postgres:securepassword123@localhost:5432/web_db?sslmode=disable"
  testDB, err := sql.Open("postgres", connStr)
  if err != nil {
    t.Fatalf("No se pudo conectar a la base de datos: %v", err)
  }
  defer testDB.Close()

  if err := testDB.Ping(); err != nil {
    t.Fatalf("La base de datos no responde: %v", err)
  }

  queries := New(testDB)
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  // crear materia
  materia, err := queries.CreateMateria(ctx, CreateMateriaParams{
    Nombre: "Redes de Computadoras I",
    Descripcion: sql.NullString{
      String: "Materia sobre arquitecturas y protocolos de red",
      Valid:  true,
    },
  })

  if err != nil {
    t.Fatalf("Error al crear materia base: %v", err)
  }

  var preguntaID int32

  // crear pregunta
  t.Run("Crear Pregunta", func(t *testing.T) {
    pregunta, err := queries.CreatePregunta(ctx, CreatePreguntaParams{
      Enunciado: "¿En qué capa del modelo OSI opera el protocolo IP?",
      IDMateria: materia.IDMateria,
    })

    if err != nil {
      t.Fatalf("Error al crear la pregunta: %v", err)
    }

    if pregunta.IDPregunta == 0 {
      t.Fatalf("Se esperaba un IDPregunta autoincremental válido")
    }

    if pregunta.IDMateria != materia.IDMateria {
      t.Errorf("IDMateria no coincide: obtenido %d, esperado %d", pregunta.IDMateria, materia.IDMateria)
    }

    preguntaID = pregunta.IDPregunta
  })

  // leer pregunta
  t.Run("Leer Pregunta", func(t *testing.T) {
    pregunta, err := queries.GetPregunta(ctx, preguntaID)

    if err != nil {
      t.Fatalf("Error al obtener la pregunta con ID %d: %v", preguntaID, err)
    }

    enunciadoEsperado := "¿En qué capa del modelo OSI opera el protocolo IP?"
    if pregunta.Enunciado != enunciadoEsperado {
      t.Errorf("Obtenido %q, se esperaba %q", pregunta.Enunciado, enunciadoEsperado)
    }
  })

  // listar preguntas x materia
  t.Run("Listar Preguntas por Materia", func(t *testing.T) {
    preguntas, err := queries.ListPreguntasByMateria(ctx, materia.IDMateria)
    if err != nil {
      t.Fatalf("Error al listar preguntas de la materia: %v", err)
    }

    if len(preguntas) == 0 {
      t.Errorf("Se esperaba al menos 1 pregunta en la lista")
    }
  })

  // actualizar pregunta
  t.Run("Actualizar Pregunta", func(t *testing.T) {
    nuevoEnunciado := "¿En qué capa del modelo TCP/IP opera el protocolo IP?"
    err := queries.UpdatePregunta(ctx, UpdatePreguntaParams{
      IDPregunta: preguntaID,
      Enunciado: nuevoEnunciado,
      IDMateria: materia.IDMateria,
    })

    if err != nil {
      t.Fatalf("Error al actualizar la pregunta: %v", err)
    }

    preguntaAct, err := queries.GetPregunta(ctx, preguntaID)
    if err != nil {
      t.Fatalf("Error al recuperar la pregunta actualizada: %v", err)
    }

    if preguntaAct.Enunciado != nuevoEnunciado {
      t.Errorf("El enunciado no se actualizó: obtenido %q, esperado %q", preguntaAct.Enunciado, nuevoEnunciado)
    }
  })

  // eliminar pregunta
  t.Run("Eliminar Pregunta", func(t *testing.T) {
    err := queries.DeletePregunta(ctx, preguntaID)
    if err != nil {
      t.Fatalf("Error al eliminar la pregunta: %v", err)
    }

    _, errQuery := queries.GetPregunta(ctx, preguntaID)
    if errQuery == nil {
      t.Errorf("La pregunta con ID %d sigue existiendo en la BD", preguntaID)
    } else if !errors.Is(errQuery, sql.ErrNoRows) {
      t.Errorf("Se esperaba sql.ErrNoRows, pero se obtuvo: %v", errQuery)
    }
  })

  // limpieza de materia base
  _ = queries.DeleteMateria(ctx, materia.IDMateria)
}
