package pantallas

import "github.com/rivo/tview"

func PantallaConfiguraciones() *tview.TextView {

	configuraciones := tview.NewTextView()

	configuraciones.SetTitle("Seccion de Configuraciones").
		SetTitleAlign(tview.AlignCenter).
		SetBorder(true)

	return configuraciones

}
