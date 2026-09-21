package tests

import (
  "context"
  "database/sql"
  "errors"
  "testing"
  "time"

  _ "github.com/lib/pq"
  . "ejemplo.com/tp-especial/db/sqlc"
)

func TestMateria_CRUD(t *testing.T) {
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

  var materiaID int32

  // create "Comunicación de datos I"
  t.Run("Crear Materia", func(t *testing.T) {
    materia, err := queries.CreateMateria(ctx, CreateMateriaParams{
      Nombre: "Comunicación de Datos I",
      Descripcion: sql.NullString{
        String: "Introducción a redes y transmisión de datos",
        Valid: true,
      },
    })

    if err != nil {
      t.Fatalf("Error al crear materia: %v", err)
    }

    if materia.IDMateria == 0 {
      t.Fatalf("Se esperaba un ID válido generado por PostgreSQL")
    }

    materiaID = materia.IDMateria
  })

  // read
  t.Run("Leer Materia", func(t *testing.T) {
    materia, err := queries.GetMateria(ctx, materiaID)
    if err != nil {
      t.Fatalf("Error al obtener materia con ID %d: %v", materiaID, err)
    }

    if materia.Nombre != "Comunicación de Datos I" {
      t.Errorf("Obtenido %q, se esperaba %q", materia.Nombre, "Comunicación de Datos I")
    }
  })

  // update "Redes de Computadoras I"
  t.Run("Actualizar Materia", func(t *testing.T) {
    nuevoNombre := "Redes de Computadoras I"
    err := queries.UpdateMateria(ctx, UpdateMateriaParams{
      IDMateria: materiaID,
      Nombre: nuevoNombre,
      Descripcion: sql.NullString{
        String: "Introducción a arquitecturas de red y protocolos",
	      Valid: true,
      },
    })

    if err != nil {
      t.Fatalf("Error al actualizar materia: %v", err)
    }

    materiaActualizada, err := queries.GetMateria(ctx, materiaID)
    if err != nil {
      t.Fatalf("Error al recuperar materia actualizada: %v", err)
    }

    if materiaActualizada.Nombre != nuevoNombre {
      t.Errorf("El nombre no se actualizó: obtenido %q, se esperaba %q", materiaActualizada.Nombre, nuevoNombre)
    }
  })

  // delete
  t.Run("Eliminar Materia", func(t *testing.T) {
    err := queries.DeleteMateria(ctx, materiaID)
    if err != nil {
      t.Fatalf("Error al eliminar materia: %v", err)
    }

    _, errQuery := queries.GetMateria(ctx, materiaID)
    if errQuery == nil {
      t.Errorf("La materia con ID %d sigue existiendo en la BD", materiaID)
    } else if !errors.Is(errQuery, sql.ErrNoRows) {
      t.Errorf("Se esperaba sql.ErrNoRows, pero se obtuvo: %v", errQuery)
    }
  })
}
