package db

import (
	"database/sql"
)

func CrearBaseDeDatos() error {

	// Crear y/o abrir la base de datos (si no existe, se crea)
	db, err := ConectarBaseDeDatos()
	if err != nil {
		return err // Devolvemos nil y el error en caso de falla
	}
	// Crear una tabla (si no existe)
	sqlStmt := `
CREATE TABLE IF NOT EXISTS productos (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	tipo TEXT,
	nombre TEXT,
	codigo TEXT,
	costo REAL,
	precio REAL,
	garantia INT,
	proveedor TEXT,
	comentario TEXT,
	stock int

);
`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err // Devolvemos nil y el error en caso de falla
	}
	// Devolvemos la conexión a la base de datos y nil para el error
	return nil
}

func ConectarBaseDeDatos() (*sql.DB, error) {

	// Crear y/o abrir la base de datos (si no existe, se crea)
	db, err := sql.Open("sqlite3", "./db/sistema.db")
	if err != nil {
		return nil, err // Devolvemos nil y el error en caso de falla
	}
	return db, nil
}
