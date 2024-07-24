package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Owbird/SVault-Engine/pkg/config"
	"github.com/owbird/svault/internal/server"
	"github.com/skratchdot/open-golang/open"
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
					if l.Error != nil {
						logsContainer.Add(
							widget.NewRichText(
								&widget.TextSegment{
									Text: fmt.Sprintf("[!] API Error: %v", l.Error),
									Style: widget.RichTextStyle{
										ColorName: "red",
									},
								},
							),
						)
					} else {
						logsContainer.Add(
							widget.NewRichText(&widget.TextSegment{Text: fmt.Sprintf("[+] API Log: %v", l.Message)}))
					}

				case "web_ui_build":
					if l.Error != nil {
						logsContainer.Add(
							widget.NewRichText(
								&widget.TextSegment{
									Text: fmt.Sprintf("[!] Web Build Error: %v", l.Error),
									Style: widget.RichTextStyle{
										ColorName: "red",
									},
								},
							),
						)
					} else {
						logsContainer.Add(
							widget.NewRichText(&widget.TextSegment{Text: fmt.Sprintf("[+] Web Build Running: %v", l.Message)}))
					}
				case "serve_web_ui_local":
					if l.Error != nil {
						logsContainer.Add(
							widget.NewRichText(
								&widget.TextSegment{
									Text: fmt.Sprintf("[!] Local Web Error: %v", l.Error),
									Style: widget.RichTextStyle{
										ColorName: "red",
									},
								},
							),
						)
					} else {
						logsContainer.Add(
							widget.NewRichText(&widget.TextSegment{Text: fmt.Sprintf("[+] Local Web Running: %v", l.Message)}))
					}
				case "serve_web_ui_remote":
					logsContainer.Add(
						widget.NewRichText(&widget.TextSegment{Text: fmt.Sprintf("[+] Remote Web Running: %v", l.Message)}))

					open.Run(l.Message)
					open.Run("https://loca.lt/mytunnelpassword")

				}

				logsContainer.Refresh()
			}
		}()
	}, sui.Window)
}

func (sui *ServerUI) ServerSettings() {

	settingsWindow := fyne.CurrentApp().NewWindow("Server settings")
	settingsWindow.Resize(fyne.NewSize(500, 500))

	appConfig := config.NewAppConfig()

	serverConfig := appConfig.GetSeverConfig()

	serverNameInput := widget.NewEntry()
	serverNameInput.SetPlaceHolder("Enter name")

	allowUploadsChecker := widget.NewCheck("Allow uploads", func(value bool) {
		serverConfig.SetAllowUploads(value)
	})

	serverNameInput.Text = serverConfig.Name
	allowUploadsChecker.Checked = serverConfig.AllowUploads

	saveBtn := widget.NewButton("Save", func() {

		serverConfig.SetName(serverNameInput.Text)

		appConfig.Save()

		settingsWindow.Close()

	})

	layoutContainer := container.NewVBox(serverNameInput, allowUploadsChecker, saveBtn)

	settingsWindow.SetContent(layoutContainer)

	settingsWindow.Show()

}
