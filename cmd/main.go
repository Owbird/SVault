package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/owbird/svault/internal/ui"
	"github.com/owbird/svault/internal/vfs"
)

func main() {
	a := app.New()

	w := a.NewWindow("SVault")

	w.Resize(fyne.NewSize(500, 500))

	hui := ui.NewHomeUI(w)
	sui := ui.NewServerUI(w)

	vfs := vfs.NewVFSFunctions()

	menus := []*fyne.Menu{
		fyne.NewMenu("Vaults", fyne.NewMenuItem("New Vault", hui.CreateVault)),
		fyne.NewMenu("File Server",
			fyne.NewMenuItem("Host", sui.ChooseHostDir),
			fyne.NewMenuItem("Settings", sui.ServerSettings),
		),
		fyne.NewMenu("VFS", fyne.NewMenuItem("Mount", vfs.Mount)),
	}

	w.SetMainMenu(fyne.NewMainMenu(menus...))

	w.SetContent(hui.Home())

	w.Show()

	a.Run()
}
