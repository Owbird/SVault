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

	hui := ui.NewHomeUI(w)
	sui := ui.NewServerUI(w)

	menus := []*fyne.Menu{
		fyne.NewMenu("Vaults", fyne.NewMenuItem("New Vault", hui.CreateVault)),
		fyne.NewMenu("File Sharing",
			fyne.NewMenuItem("Host Directory", sui.ChooseHostDir),
			fyne.NewMenuItem("Share file", sui.ShareFile),
			fyne.NewMenuItem("Receive file", sui.ReceiveFile),
			fyne.NewMenuItem("Settings", sui.ServerSettings),
		),
	}

	w.SetMainMenu(fyne.NewMainMenu(menus...))

	w.SetContent(hui.Home())

	w.Show()

	a.Run()
}
