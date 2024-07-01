package ui

import (
	"fmt"
	"path/filepath"

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

func (hui *HomeUI) ViewVault(vault string) {
	vaultPwdInput := widget.NewEntry()
	vaultPwdInput.SetPlaceHolder("Enter Password")
	vaultPwdInput.Validator = validation.NewRegexp("\\S+", "Field is required")

	vaultPwdForm := widget.NewFormItem("Vault password", vaultPwdInput)

	formItems := []*widget.FormItem{
		vaultPwdForm,
	}

	callback := func(proceed bool) {
		if proceed {

			pwdMatch, err := hui.Vault.AuthVault(vault, vaultPwdInput.Text)
			if err != nil {
				dialog.NewError(err, hui.Window).Show()
				return
			}

			if !pwdMatch {
				dialog.NewError(fmt.Errorf("Passwords do not match"), hui.Window).Show()
				return
			}

			vaultWindow := fyne.CurrentApp().NewWindow(fmt.Sprintf("%v vault", vault))

			updateContentUI := func() {
				fileCards := []fyne.CanvasObject{}

				files, err := hui.Vault.GetVault(vault, vaultPwdInput.Text)
				if err != nil {
					dialog.NewError(err, hui.Window).Show()
					return
				}

				for _, file := range files {
					fileCards = append(fileCards, widget.NewCard(
						filepath.Base(file.Name),
						file.ModTime.Format("2nd January, 2006"),
						container.NewStack(),
					),
					)
				}

				vaultWindow.SetContent(container.NewHBox(fileCards...))
			}

			menus := []*fyne.Menu{
				fyne.NewMenu("File", fyne.NewMenuItem("Add file", func() {
					dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
						if err != nil {
							dialog.NewError(err, vaultWindow).Show()
						}

						hui.Vault.AddFile(uc.URI().Path(), vault, vaultPwdInput.Text)

						updateContentUI()
					}, vaultWindow)
				})),
			}

			vaultWindow.SetMainMenu(fyne.NewMainMenu(menus...))

			updateContentUI()

			vaultWindow.Resize(fyne.NewSize(500, 500))
			vaultWindow.Show()
		}
	}

	formDialog := dialog.NewForm("Authorize Vault",
		"Proceed",
		"Cancel",
		formItems,
		callback,
		hui.Window,
	)

	formDialog.Resize(fyne.NewSize(300, 200))
	formDialog.Show()
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
