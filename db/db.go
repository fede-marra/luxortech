package db

import (
	"database/sql"
	"log"
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

func ObtenerListaProductos(db *sql.DB) []string {
	rows, err := db.Query("SELECT tipo FROM tiposProductos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tipos []string
	for rows.Next() {
		var tipo string
		if err := rows.Scan(&tipo); err != nil {
			log.Fatal(err)
		}
		tipos = append(tipos, tipo)
	}
	return tipos
}
