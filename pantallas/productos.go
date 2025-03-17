package pantallas

import (
	"fmt"
	"sistema/db"
	"strconv"

	"github.com/rivo/tview"
)

func PantallaProductos() tview.Primitive {

	productos := tview.NewTextView()
	productos.SetTitle("Seccion de Productos").
		SetTitleAlign(tview.AlignCenter).
		SetBorder(true)

	newProducto := tview.NewForm().
		AddInputField("Tipo", "", 20, nil, nil).
		AddInputField("Nombre", "", 20, nil, nil).
		AddInputField("Codigo", "", 20, nil, nil).
		AddInputField("Costo", "", 20, nil, nil).
		AddInputField("Precio", "", 20, nil, nil).
		AddInputField("Garantia", "", 20, nil, nil).
		AddInputField("Proveedor", "", 20, nil, nil).
		AddInputField("Comentario", "", 20, nil, nil).
		AddInputField("Stock", "", 20, nil, nil)

	newProducto.AddButton("Guardar", func() {
		tipo := newProducto.GetFormItemByLabel("Tipo").(*tview.InputField).GetText()
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
			fmt.Println("Error, Costo Invalido")
			return
		}

		precio, err := strconv.ParseFloat(precioStr, 64)
		if err != nil {
			fmt.Println("Error, Precio Invalido")
			return
		}

		stock, err := strconv.ParseFloat(stockStr, 64)
		if err != nil {
			fmt.Println("Error, stock Invalido")
			return
		}
		// Guardar en la base de datos

		err = db.InsertarProducto(tipo, nombre, codigo, costo, precio, garantia, proveedor, comentario,int(stock))
		if err != nil {
			fmt.Println("Error al guardar el producto:", err)
		}

	}).
		AddButton("Cancelar", func() {
			fmt.Println("Operacion Cancelada")

		}).
		SetBorder(true).SetTitle("Formulario de Productos")

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(productos, 3, 1, false).
		AddItem(newProducto, 0, 3, true)
	return layout

}
