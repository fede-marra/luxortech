package pantallas

import (
	"fmt"
	"sistema/db"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func PantallaProductos(pages *tview.Pages) tview.Primitive {

	/*productos := tview.NewTextView()
	productos.SetTitle("Seccion de Productos").
		SetTitleAlign(tview.AlignCenter).
		SetBorder(true)
	*/
	base, _ := db.ConectarBaseDeDatos()
	tipos := db.ObtenerListaProductos(base)

	tabla := db.CargarTablaProductos()


	dropdown := tview.NewDropDown().
		SetLabel("Tipos").
		SetOptions(tipos, nil)
	newProducto := tview.NewForm().
		AddFormItem(dropdown).
		AddInputField("Nombre", "", 20, nil, nil).
		AddInputField("Codigo", "", 20, nil, nil).
		AddInputField("Costo", "", 20, nil, nil).
		AddInputField("Precio", "", 20, nil, nil).
		AddInputField("Garantia", "", 20, nil, nil).
		AddInputField("Proveedor", "", 20, nil, nil).
		AddInputField("Comentario", "", 20, nil, nil).
		AddInputField("Stock", "", 20, nil, nil)

	newProducto.SetTitle("Formulario").SetBorder(true)
	newProducto.AddButton("Guardar", func() {
		_, tipo := dropdown.GetCurrentOption()
		nombre := newProducto.GetFormItemByLabel("Nombre").(*tview.InputField).GetText()
		codigo := newProducto.GetFormItemByLabel("Codigo").(*tview.InputField).GetText()
		costoStr := newProducto.GetFormItemByLabel("Costo").(*tview.InputField).GetText()
		precioStr := newProducto.GetFormItemByLabel("Precio").(*tview.InputField).GetText()
		garantia := newProducto.GetFormItemByLabel("Garantia").(*tview.InputField).GetText()
		proveedor := newProducto.GetFormItemByLabel("Proveedor").(*tview.InputField).GetText()
		comentario := newProducto.GetFormItemByLabel("Comentario").(*tview.InputField).GetText()
		stockStr := newProducto.GetFormItemByLabel("Stock").(*tview.InputField).GetText()

		// Convertir strings a tipos numéricos

		costo, err := strconv.ParseFloat(costoStr, 64)
		if err != nil {
			fmt.Printf("Error, Costo Invalido: %v\n", err)
			return
		}

		precio, err := strconv.ParseFloat(precioStr, 64)
		if err != nil {
			fmt.Printf("Error, Precio Invalido: %v\n", err)
			return
		}

		stock, err := strconv.ParseFloat(stockStr, 64)
		if err != nil {
			fmt.Printf("Error, stock Invalido: %v\n", err)
			return
		}

		// Guardar en la base de datos

		err = db.InsertarProducto(tipo, nombre, codigo, costo, precio, garantia, proveedor, comentario, int(stock))
		if err != nil {
			fmt.Printf("Error al guardar el producto: %v\n", err)
		}
		//Actualiza la tabla de productos

		db.ActualizarTabla(tabla)

		// Función para mostrar el modal de confirmación
		mostrarAviso := func(mensaje string) {
			modal := tview.NewModal().
				SetText(mensaje).
				AddButtons([]string{"OK"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					pages.RemovePage("aviso") // Cerrar modal al presionar "OK"
				})
			modal.SetBackgroundColor(tcell.ColorDarkGreen).SetBorder(false)
			pages.AddPage("aviso", modal, true, true)
		}
		mostrarAviso("Producto Guardado Correctamente")
		dropdown.SetCurrentOption(-1)
		newProducto.GetFormItemByLabel("Nombre").(*tview.InputField).SetText("")
		newProducto.GetFormItemByLabel("Codigo").(*tview.InputField).SetText("")
		newProducto.GetFormItemByLabel("Costo").(*tview.InputField).SetText("")
		newProducto.GetFormItemByLabel("Precio").(*tview.InputField).SetText("")
		newProducto.GetFormItemByLabel("Garantia").(*tview.InputField).SetText("")
		newProducto.GetFormItemByLabel("Proveedor").(*tview.InputField).SetText("")
		newProducto.GetFormItemByLabel("Comentario").(*tview.InputField).SetText("")
		newProducto.GetFormItemByLabel("Stock").(*tview.InputField).SetText("")
	}).
		AddButton("Cancelar", func() {
			dropdown.SetCurrentOption(-1)
			newProducto.GetFormItemByLabel("Nombre").(*tview.InputField).SetText("")
			newProducto.GetFormItemByLabel("Codigo").(*tview.InputField).SetText("")
			newProducto.GetFormItemByLabel("Costo").(*tview.InputField).SetText("")
			newProducto.GetFormItemByLabel("Precio").(*tview.InputField).SetText("")
			newProducto.GetFormItemByLabel("Garantia").(*tview.InputField).SetText("")
			newProducto.GetFormItemByLabel("Proveedor").(*tview.InputField).SetText("")
			newProducto.GetFormItemByLabel("Comentario").(*tview.InputField).SetText("")
			newProducto.GetFormItemByLabel("Stock").(*tview.InputField).SetText("")
		}).
		SetBorder(true).SetTitle("Formulario de Productos")
	

	


	layout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(newProducto, 0, 1, true).
		AddItem(tabla, 0, 1, true)
	return layout

}
