package pantallas

import "github.com/rivo/tview"

func MostrarMensaje(app *tview.Application, pages *tview.Pages, mensaje string) {
	modal := tview.NewModal().
		SetText(mensaje).
		AddButtons([]string{"Aceptar"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.RemovePage("mensaje") // Cierra la ventana al presionar "Aceptar"
		})

	pages.AddPage("mensaje", modal, true, true)
}
