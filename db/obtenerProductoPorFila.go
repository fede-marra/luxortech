package db

import (
	"database/sql"
)


// ObtenerProductoPorFila obtiene un producto basado en su posición en la tabla visual
func ObtenerProductoPorFila(index int) Producto {
	db, err := ConectarBaseDeDatos()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "SELECT id, nombre, codigo, costo, precio, garantia, proveedor, comentario, stock FROM productos ORDER BY id LIMIT 1 OFFSET ?"
	row := db.QueryRow(query, index)

	var p Producto
	err = row.Scan(&p.Id, &p.Nombre, &p.Codigo, &p.Costo, &p.Precio, &p.Garantia, &p.Proveedor, &p.Comentario, &p.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			return Producto{} // Si no encuentra nada, devuelve un producto vacío
		}
		panic(err)
	}

	return p
}
