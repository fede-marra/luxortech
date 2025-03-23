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

	pages := tview.NewPages()
	// Agregar pantallas desde los archivos separados
	pages.AddPage("Ventas", pantallas.PantallaVentas(), true, true)
	pages.AddPage("Productos", pantallas.PantallaProductos(pages), true, false)
	pages.AddPage("Gastos", pantallas.PantallaGastos(), true, false)
	pages.AddPage("Clientes", pantallas.PantallaClientes(), true, false)
	pages.AddPage("Configuraciones", pantallas.PantallaConfiguraciones(), true, false)

	menu := tview.NewList().
		AddItem("Ventas", "", '1', func() { pages.SwitchToPage("Ventas") }).
		AddItem("Productos", "", '2', func() { pages.SwitchToPage("Productos").SendToFront("Productos") }).
		AddItem("Gastos", "", '3', func() { pages.SwitchToPage("Gastos") }).
		AddItem("Clientes", "", '4', func() { pages.SwitchToPage("Clientes") }).
		AddItem("Configuraciones", "", '5', func() { pages.SwitchToPage("Configuraciones") }).
		AddItem("Salir", "", '6', func() {
			app.Stop()
		})
	menu.SetBorder(true).SetTitle("Menu")

	flex := pantallas.MostrarLayout(menu, pages)

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}

}
