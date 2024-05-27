package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/owbird/svault/internal/ui"
)

func main() {
	a := app.New()

	w := a.NewWindow("SVault")

	w.Resize(fyne.NewSize(500, 500))

	uf := ui.NewUIFunctions(w)

	menus := []*fyne.Menu{
		fyne.NewMenu("Vaults", fyne.NewMenuItem("New Vault", uf.CreateVault)),
	}

	w.SetMainMenu(fyne.NewMainMenu(menus...))

	w.SetContent(uf.Home())

	w.Show()

	a.Run()
}
