package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var menus = []*fyne.Menu{
	fyne.NewMenu("Vaults", fyne.NewMenuItem("New", func() {})),
}

func main() {
	a := app.New()

	w := a.NewWindow("SVault")

	w.Resize(fyne.NewSize(500, 500))

	w.SetMainMenu(fyne.NewMainMenu(menus...))

	w.Show()

	a.Run()
}
