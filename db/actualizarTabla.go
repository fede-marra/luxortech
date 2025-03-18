package db

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ActualizarTabla(tabla *tview.Table) {

	base, _ := ConectarBaseDeDatos()
	productos, err := ObtenerTodosLosProductos(base)
	if err != nil {
		fmt.Printf("Error al obtener productos: %v\n", err)
		return
	}

	// Limpiar la tabla antes de volver a cargar los datos
	tabla.Clear()

	// Volver a agregar los encabezados
	headers := []string{"Id", "Tipo", "Nombre", "Codigo", "Precio", "Garantia", "Proveedor", "Comentario", "Stock"}
	for i, header := range headers {
		tabla.SetCell(0, i, tview.NewTableCell(header).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignCenter).
			SetSelectable(false))
	}

	// Volver a agregar los productos
	for i, prod := range productos {
		tabla.SetCell(i+1, 0, tview.NewTableCell(fmt.Sprintf("%d", prod.Id)).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignCenter).
			SetSelectable(true))
		tabla.SetCell(i+1, 1, tview.NewTableCell(prod.Tipo).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignCenter).
			SetSelectable(true))
		tabla.SetCell(i+1, 2, tview.NewTableCell(prod.Nombre).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignCenter).
			SetSelectable(true))
		tabla.SetCell(i+1, 3, tview.NewTableCell(prod.Codigo).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignCenter).
			SetSelectable(true))
		tabla.SetCell(i+1, 4, tview.NewTableCell(fmt.Sprintf("%.2f", prod.Precio)).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignCenter).
			SetSelectable(true))
		tabla.SetCell(i+1, 5, tview.NewTableCell(prod.Garantia).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignCenter).
			SetSelectable(true))
		tabla.SetCell(i+1, 6, tview.NewTableCell(prod.Proveedor).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignCenter).
			SetSelectable(true))
		tabla.SetCell(i+1, 7, tview.NewTableCell(prod.Comentario).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignCenter).
			SetSelectable(true))
		tabla.SetCell(i+1, 8, tview.NewTableCell(fmt.Sprintf("%d", prod.Stock)).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignCenter).
			SetSelectable(true))
	}
}
