package ui

import (
	"fmt"
	"path/filepath"

	"fyne.io/fyne/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/owbird/svault/internal/vault"
)

type HomeUI struct {
	Window fyne.Window
	Vault  *vault.VaultFunctions
}

func NewHomeUI(window fyne.Window) *HomeUI {
	return &HomeUI{
		Vault:  vault.NewVaultFunctions(),
		Window: window,
	}
}

func (hui *HomeUI) AuthorizeVault(vault string, callback func(pwd string), window fyne.Window) {
	vaultPwdInput := widget.NewEntry()
	vaultPwdInput.SetPlaceHolder("Enter Password")
	vaultPwdInput.Validator = validation.NewRegexp("\\S+", "Field is required")

	vaultPwdForm := widget.NewFormItem("Vault password", vaultPwdInput)

	formItems := []*widget.FormItem{
		vaultPwdForm,
	}

	formDialog := dialog.NewForm("Authorize Vault",
		"Proceed",
		"Cancel",
		formItems,
		func(proceed bool) {
			if proceed {
				 err := hui.Vault.AuthVault(vault, vaultPwdInput.Text)
				if err != nil {
					dialog.NewError(err, window).Show()
					return
				}

				callback(vaultPwdInput.Text)
			}
		},
		window,
	)

	formDialog.Resize(fyne.NewSize(300, 200))
	formDialog.Show()
}

func (hui *HomeUI) ViewVault(vault string) {
	callback := func(pwd string) {
		vaultWindow := fyne.CurrentApp().NewWindow(fmt.Sprintf("%v vault", vault))

		files, err := hui.Vault.GetVault(vault, pwd)
		if err != nil {
			dialog.NewError(err, hui.Window).Show()
			return
		}

		fileList := widget.NewList(
			func() int { return len(files) },
			func() fyne.CanvasObject {
				return container.NewHBox(widget.NewIcon(theme.FileIcon()), widget.NewLabel(""))
			},
			func(id widget.ListItemID, item fyne.CanvasObject) {
				hBox := item.(*fyne.Container)
				hBox.Objects[1].(*widget.Label).SetText(filepath.Base(files[id].Name))
			},
		)

		updateFiles := func() {
			files, err = hui.Vault.GetVault(vault, pwd)
			if err != nil {
				dialog.NewError(err, hui.Window).Show()
				return
			}

			fileList.Refresh()
		}

		fileList.OnSelected = func(id widget.ListItemID) {
			menu := fyne.NewMenu("Actions",
				fyne.NewMenuItem("Delete", func() {
					hui.AuthorizeVault(vault, func(pwd string) {
						hui.Vault.DeleteFile(files[id].Name, vault, pwd)

						updateFiles()
					}, vaultWindow)
				}),
			)
			widget.ShowPopUpMenuAtPosition(menu, vaultWindow.Canvas(), fyne.CurrentApp().Driver().AbsolutePositionForObject(fileList))
		}

		vaultWindow.SetContent(fileList)

		menus := []*fyne.Menu{
			fyne.NewMenu("File", fyne.NewMenuItem("Add file", func() {
				dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
					if err != nil {
						dialog.NewError(err, vaultWindow).Show()
					}

					if uc != nil {
						hui.Vault.AddFile(uc.URI().Path(), vault, pwd)
					}

					updateFiles()
				}, vaultWindow)
			})),
		}

		vaultWindow.SetMainMenu(fyne.NewMainMenu(menus...))

		vaultWindow.Resize(fyne.NewSize(500, 500))
		vaultWindow.Show()
	}

	hui.AuthorizeVault(vault, callback, hui.Window)
}

func (hui *HomeUI) Home() fyne.CanvasObject {
	vaults, err := hui.Vault.ListVaults()
	if err != nil {
		dialog.NewError(err, hui.Window).Show()
	}

	if len(vaults) == 0 {
		return container.NewCenter(widget.NewRichTextWithText("No vaults yet"))
	}

	cards := []fyne.CanvasObject{}

	for _, vault := range vaults {
		button := widget.NewButton(vault.Name, func() { hui.ViewVault(vault.Name) })

		card := container.NewCenter(button)
		cards = append(cards, card)

	}

	return container.NewGridWithColumns(4, cards...)
}

func (hui *HomeUI) CreateVault() {
	vaultNameInput := widget.NewEntry()
	vaultNameInput.SetPlaceHolder("Enter name")
	vaultNameInput.Validator = validation.NewRegexp("\\S+", "Field is required")

	vaultPwdInput := widget.NewEntry()
	vaultPwdInput.SetPlaceHolder("Enter Password")
	vaultPwdInput.Validator = validation.NewRegexp("\\S+", "Field is required")

	vaultNameForm := widget.NewFormItem("Vault name", vaultNameInput)
	vaultPwdForm := widget.NewFormItem("Vault password", vaultPwdInput)

	formItems := []*widget.FormItem{
		vaultNameForm,
		vaultPwdForm,
	}

	callback := func(create bool) {
		if create {
			err := hui.Vault.CreateVault(vaultNameInput.Text, vaultPwdInput.Text)
			if err != nil {
				dialog.NewError(err, hui.Window).Show()
			}

			hui.Window.SetContent(hui.Home())
		}
	}

	formDialog := dialog.NewForm("Create New Vault",
		"Create",
		"Cancel",
		formItems,
		callback,
		hui.Window,
	)

	formDialog.Resize(fyne.NewSize(300, 200))
	formDialog.Show()
}
