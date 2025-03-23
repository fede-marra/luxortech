package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // O el driver que uses (mysql, postgres, etc.)
)

// Producto representa un producto en la base de datos
type Producto struct {
	Id         int
	Tipo       string
	Nombre     string
	Codigo     string
	Precio     float64
	Costo      float64
	Garantia   string
	Proveedor  string
	Comentario string
	Stock      int
}

// ObtenerTodosLosProductos obtiene todos los productos de la base de datos
func ObtenerTodosLosProductos(db *sql.DB, texto string) ([]Producto, error) {
	rows, err := db.Query("SELECT id, tipo, nombre, codigo, precio, garantia, proveedor, comentario, stock FROM productos WHERE nombre LIKE ?", texto+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []Producto
	for rows.Next() {
		var p Producto
		err := rows.Scan(&p.Id, &p.Tipo, &p.Nombre, &p.Codigo, &p.Precio, &p.Garantia, &p.Proveedor, &p.Comentario, &p.Stock)
		if err != nil {
			return nil, err
		}
		productos = append(productos, p)
	}

	return productos, nil
}
