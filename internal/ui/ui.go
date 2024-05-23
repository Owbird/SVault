package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/owbird/svault/internal/vault"
)

type UIFunctions struct {
	Window fyne.Window
	Vault  *vault.VaultFunctions
}

func NewUIFunctions(window fyne.Window) *UIFunctions {
	return &UIFunctions{
		Vault:  vault.NewVaultFunctions(),
		Window: window,
	}
}

func (uf *UIFunctions) CreateVault() {

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
			err := uf.Vault.CreateVault(vaultNameInput.Text, vaultPwdInput.Text)

			if err != nil {
				dialog.NewError(err, uf.Window).Show()
			}
		}
	}

	formDialog := dialog.NewForm("Create New Vault",
		"Create",
		"Cancel",
		formItems,
		callback,
		uf.Window,
	)

	formDialog.Resize(fyne.NewSize(300, 200))
	formDialog.Show()
}
