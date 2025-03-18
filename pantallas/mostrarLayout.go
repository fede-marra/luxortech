package pantallas

import "github.com/rivo/tview"

func MostrarLayout(menus tview.Primitive, paginas tview.Primitive) tview.Primitive {
	flex := tview.NewFlex().
		AddItem(menus, 0, 1, true).
		AddItem(paginas, 0, 6, false)

	return flex
}
