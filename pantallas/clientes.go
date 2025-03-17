package pantallas

import (
	"github.com/rivo/tview"
)

func PantallaClientes() *tview.TextView {

	clientes := tview.NewTextView()
	clientes.SetTitle("Seccion de Clientes").
	SetTitleAlign(tview.AlignCenter).
	SetBorder(true)

	return clientes
}
