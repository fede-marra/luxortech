package pantallas

import (
	"fmt"
	"sistema/db"
	"strconv"

	"github.com/rivo/tview"
)

func PantallaProductos() tview.Primitive {

	productos := tview.NewBox()
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

		err = db.InsertarProducto(tipo, nombre, codigo, costo, precio, garantia, proveedor, comentario,int(stock))
		if err != nil {
			fmt.Printf("Error al guardar el producto: %v\n", err)
		}

	}).
		AddButton("Cancelar", func() {
			fmt.Println("Operacion Cancelada")

		}).
		SetBorder(true).SetTitle("Formulario de Productos")

		


	//productos.SetText(newProducto.GetTitle())
	return productos

}
