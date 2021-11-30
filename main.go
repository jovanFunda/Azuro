package main

import (
	_ "embed"
	"log"

	"github.com/wailsapp/wails"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

	customers, err := NewCustomers()
	if err != nil {
		log.Fatal(err)
	}

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1080,
		Height:    720,
		Title:     "Tepih servis Azuro",
		Resizable: false,
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
	})

	app.Bind(customers)
	app.Run()
}
