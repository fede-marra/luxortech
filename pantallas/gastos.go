package pantallas

import "github.com/rivo/tview"

func PantallaGastos() *tview.TextView {

	gastos := tview.NewTextView()
	gastos.SetTitle("Seccion de Gastos").
	SetTitleAlign(tview.AlignCenter).
	SetBorder(true)

	return gastos

}
