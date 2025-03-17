package main

import (
	"fmt"
	"log"
	"sistema/db"
	"sistema/pantallas"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	app.EnableMouse(true)

	// Llamamos a la función CrearBaseDeDatos desde el paquete db
	err := db.CrearBaseDeDatos()
	if err != nil {
		log.Fatal(err)
	}
	dbConn, err := db.ConectarBaseDeDatos()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// Aquí puedes continuar con el uso de la base de datos
	fmt.Println("Base de datos y tabla creadas exitosamente.")

	// Contenedor de Pantallas
	/*

		pages := tview.NewPages()

		// Agregar pantallas desde los archivos separados
		pages.AddPage("Ventas", pantallas.PantallaVentas(), true, true)
		pages.AddPage("Productos", pantallas.PantallaProductos(), true, false)
		pages.AddPage("Gastos", pantallas.PantallaGastos(), true, false)
		pages.AddPage("Clientes", pantallas.PantallaClientes(), true, false)
		pages.AddPage("Configuraciones", pantallas.PantallaConfiguraciones(), true, false)

		menu := tview.NewList().
			AddItem("Ventas", "", '1', func() { pages.SwitchToPage("Ventas") }).
			AddItem("Productos", "", '2', func() { pages.SwitchToPage("Productos") }).
			AddItem("Gastos", "", '3', func() { pages.SwitchToPage("Gastos") }).
			AddItem("Clientes", "", '4', func() { pages.SwitchToPage("Clientes") }).
			AddItem("Configuraciones", "", '5', func() { pages.SwitchToPage("Configuraciones") }).
			AddItem("Salir", "", '6', func() {
				modal := tview.NewModal()
				modal.SetText("¿Estás seguro de que quieres salir?").
					AddButtons([]string{"Sí", "No"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						if buttonLabel == "Sí" {
							app.Stop() // Cierra la app
						} else {
							pages.SwitchToPage("Ventas") // Vuelve a la última pantalla
						}
					}).SetFocus(1)
				pages.AddPage("confirmacion", modal, true, true)
			})

		menu.SetBorder(true).
			SetTitle("Menu")
	*/

	pages := tview.NewPages()

	pages.AddPage("Productos", pantallas.PantallaProductos(), true, false)

	menu := tview.NewList().
		AddItem("Productos", "", '2', func() { pages.SwitchToPage("Productos").SendToFront("Productos") }).
		AddItem("Salir", "", '6', func() {
			app.Stop()
		})
	menu.SetBorder(true).SetTitle("Menu")

	//top := tview.NewBox().SetBorder(true).SetTitle("Top")

	//bottom := tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)")
	flex := tview.NewFlex().
		AddItem(menu, 0, 1, true).
		AddItem(pages, 0, 6, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}

}
