package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
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
		go sf.Start()

		logWindow := fyne.CurrentApp().NewWindow("Server Logs")
		logWindow.Resize(fyne.NewSize(500, 500))

		logsContainer := container.NewVBox()

		logWindow.SetContent(logsContainer)

		logWindow.Show()

		go func() {
			for l := range sf.LogCh {
				switch l.Type {
				case "api_log":
					logsContainer.Add(
						widget.NewRichText(&widget.TextSegment{Text: fmt.Sprintf("[+] API Log: %v", l.Message)}))

				case "serve_web_ui_local":
					logsContainer.Add(
						widget.NewRichText(&widget.TextSegment{Text: fmt.Sprintf("[+] Local Web Running: %v", l.Message)}))

				case "serve_web_ui_remote":
					logsContainer.Add(
						widget.NewRichText(&widget.TextSegment{Text: fmt.Sprintf("[+] Remote Web Running: %v", l.Message)}))

				}

				logsContainer.Refresh()
			}
		}()
	}, sui.Window)
}
