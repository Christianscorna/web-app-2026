package db

import (
  "context"
  "database/sql"
  "testing"
  "time"

_ "github.com/lib/pq"
)

func TestOpciones_CascadeDelete(t *testing.T) {
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

  // creamos una materia
  materia, err := queries.CreateMateria(ctx, CreateMateriaParams{
    Nombre: "Sistemas Operativos",
    Descripcion: sql.NullString{
      String: "Conceptos de kernel, procesos y memoria",
      Valid:  true,
    },
  })

  if err != nil {
    t.Fatalf("Error al crear la materia: %v", err)
  }

  // creamos una pregunta y la asociamos a la materia
  pregunta, err := queries.CreatePregunta(ctx, CreatePreguntaParams{
    Enunciado: "¿Cuál de las siguientes no es una llamada al sistema de gestión de procesos en Linux?",
    IDMateria: materia.IDMateria,
  })

  if err != nil {
    t.Fatalf("Error al crear la pregunta: %v", err)
  }

  // aca creamos una serie de opciones
  opcionesEntrada := []struct {
    texto string
    esCorrecta bool
  }{
    {"fork()", false},
    {"exec()", false},
    {"malloc()", true},
    {"wait()", false},
  }

  var IDsOpciones []int32

  t.Run("Crear 4 Opciones", func(t *testing.T) {
    for _, op := range opcionesEntrada {
      opcionGenerada, err := queries.CreateOpcion(ctx, CreateOpcionParams{
        Texto: op.texto,
	EsCorrecta: op.esCorrecta,
	IDPregunta: pregunta.IDPregunta,
      })

      if err != nil {
	t.Fatalf("Error al crear la opción %q: %v", op.texto, err)
      }

      if opcionGenerada.IDOpcion == 0 {
        t.Fatalf("Se esperaba un IDOpcion autoincremental válido")
      }

      IDsOpciones = append(IDsOpciones, opcionGenerada.IDOpcion)
    }

    if len(IDsOpciones) != 4 {
      t.Fatalf("Se esperaban 4 opciones creadas, se crearon %d", len(IDsOpciones))
    }
  })

  // verificar que se listen las opciones de la pregunta
  t.Run("Listar 4 Opciones de la Pregunta", func(t *testing.T) {
    opciones, err := queries.ListOpcionesByPregunta(ctx, pregunta.IDPregunta)

    if err != nil {
      t.Fatalf("Error al listar opciones: %v", err)
    }

    // si no son 4, algo no piolaba como tenia que piolar
    if len(opciones) != 4 {
      t.Errorf("Se esperaban 4 opciones asociadas a la pregunta, se obtuvieron %d", len(opciones))
    }
  })

  // chequeamos el borrado en cascada
  t.Run("Chequeo de delete cascade", func(t *testing.T) {
    // Al eliminar la materia, por las testricciones impuestas en el esquema,
    // se deberían borrar tambien las preguntas y las opciones asociadas
    err := queries.DeleteMateria(ctx, materia.IDMateria)
    if err != nil {
      t.Fatalf("Error al eliminar la materia padre: %v", err)
    }

    // si intentanmos listarlas ahora debieran dar error
    opcionesRestantes, err := queries.ListOpcionesByPregunta(ctx, pregunta.IDPregunta)
    if err != nil {
      t.Fatalf("Error al consultar opciones tras borrado en cascada: %v", err)
    }

    // nuevamente, la cantidad de opciones resultantes debiera ser cero
    if len(opcionesRestantes) != 0 {
      t.Errorf("Falló el borrado en cascada: la pregunta eliminada aún conserva %d opciones en la BD", len(opcionesRestantes))
    }
  })
}
