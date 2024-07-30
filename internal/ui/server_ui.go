package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Owbird/SVault-Engine/pkg/config"
	"github.com/Owbird/SVault-Engine/pkg/models"
	"github.com/owbird/svault/internal/server"
	"github.com/skratchdot/open-golang/open"
)

type ServerUI struct {
	Window    fyne.Window
	Functions *server.ServerFunctions
}

func NewServerUI(window fyne.Window) *ServerUI {
	return &ServerUI{
		Window:    window,
		Functions: server.NewServerFunctions(),
	}
}

func (sui *ServerUI) ShareFile() {
	dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.NewError(err, sui.Window)
			return
		}

		file := uc.URI().Path()

		sui.Functions.Share(file, server.ShareCallBacks{
			OnFileSent: func() {
				dialog.NewInformation("File sent", "File sent successfully", sui.Window).Show()
			},
			OnSendErr: func(err error) {
				dialog.NewError(err, sui.Window).Show()
			},
			OnProgressChange: func(progress models.FileShareProgress) {},
			OnCodeReceive: func(code string) {
				dialog.NewInformation("Code received", code, sui.Window).Show()
			},
		})

	}, sui.Window)
}

func (sui *ServerUI) ChooseHostDir() {
	dialog.ShowFolderOpen(func(lu fyne.ListableURI, err error) {
		if err != nil {
			dialog.NewError(err, sui.Window)
			return
		}

		go sui.Functions.Host(lu.Path())

		logWindow := fyne.CurrentApp().NewWindow("Server Logs")
		logWindow.Resize(fyne.NewSize(500, 500))

		logsContainer := container.NewVBox()

		logWindow.SetContent(container.NewVScroll(logsContainer))

		logWindow.Show()

		go func() {
			for l := range sui.Functions.LogCh {
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

				default:
					if l.Error != nil {
						logsContainer.Add(
							widget.NewRichText(
								&widget.TextSegment{
									Text: fmt.Sprintf("[!] Error: %v", l.Error),
								},
							),
						)
					} else {
						logsContainer.Add(
							widget.NewRichText(&widget.TextSegment{Text: fmt.Sprintf("[+] Log: %v", l.Message)}))
					}

				}
			}
			logsContainer.Refresh()
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

	serverNameInput.Text = serverConfig.GetName()
	allowUploadsChecker.Checked = serverConfig.GetAllowUploads()

	saveBtn := widget.NewButton("Save", func() {
		serverConfig.SetName(serverNameInput.Text)

		appConfig.Save()

		settingsWindow.Close()
	})

	layoutContainer := container.NewVBox(serverNameInput, allowUploadsChecker, saveBtn)

	settingsWindow.SetContent(layoutContainer)

	settingsWindow.Show()
}
