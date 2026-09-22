package coneccionbd

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	// Por ahora dejar asi luego migrar a un archivo de configuración!!!!
	/*
	   dsn := os.Getenv("DATABASE_URL")
	   if dsn == "" {
	     return nil, fmt.Errorf("DATABASE_URL is not set")
	   }
	*/
	connectionSTR := "postgres://postgres:securepassword123@localhost:5432/web_db?sslmode=disable"
	// Abrir una conexión real
	conn, err := sql.Open("postgres", connectionSTR)
	// Si no hay errores, conectarse
	if err != nil {
		fmt.Printf("Error al conectar a la base de datos: %s\n", err)
		return nil, err
	}
	// tira un ping para saber si esta viva la conexión
	err = conn.Ping()
	if err != nil {
		fmt.Printf("Error al hacer ping a la base de datos: %s\n", err)
		return nil, err
	}

	fmt.Println("Conexión exitosa")
	return conn, nil
}
