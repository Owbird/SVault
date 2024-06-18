package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/owbird/svault/internal/server"
)

type ServerUI struct {
	Window fyne.Window
}

func NewServerUI(window fyne.Window) *ServerUI {
	return &ServerUI{
		Window: window,
	}
}

func (sui *ServerUI) ChooseHostDir() {
	dialog.ShowFolderOpen(func(lu fyne.ListableURI, err error) {
		if err != nil {
			dialog.NewError(err, sui.Window)
			return
		}

		sf := server.NewServerFunctions(lu.Path())
		sf.Start()
	}, sui.Window)
}
