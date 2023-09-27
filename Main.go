package main

import (
	"demo/views"
	"fyne.io/fyne/v2"
)

func main() {
	window := views.NewMainWindow("daimaku", fyne.Size{Width: 1366, Height: 768})
	window.Start()
}
