package db

import (
	"fmt"
)

// InsertarProducto guarda un producto en la base de datos
func InsertarProducto(tipo string, nombre string, codigo string, costo float64, precio float64, garantia string, proveedor string, comentario string, stock int) error {
	db, err := ConectarBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO productos (tipo, nombre, codigo, costo, precio, garantia, proveedor, comentario, stock) VALUES (?, ?, ?)`
	_, err = db.Exec(query, tipo, nombre, codigo, costo, precio, garantia, proveedor, comentario, stock)
	if err != nil {
		return err
	}

	fmt.Println("Producto insertado correctamente")
	return nil
}
