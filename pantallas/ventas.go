package pantallas

import "github.com/rivo/tview"

// NuevaPantallaVentas crea la pantalla de ventas
func PantallaVentas() *tview.TextView {
	ventas := tview.NewTextView()
	ventas.SetTitle("Seccion Ventas").
		SetBorder(true).
		SetTitleAlign(tview.AlignCenter)
	return ventas
}
