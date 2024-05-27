package vault

import (
	"github.com/Owbird/SVault-Engine/pkg/vault"
)

type Vault struct{}

type Dir struct {
	Name    string `json:"name"`
	IsVault bool   `json:"isVault"`
	Path    string `json:"path"`
}

type VaultFunctions struct{}

var v = vault.NewVault()

func NewVaultFunctions() *VaultFunctions {
	return &VaultFunctions{}
}

func (vf *VaultFunctions) CreateVault(name, pwd string) error {
	err := v.Create(name, pwd)

	if err != nil {
		return err
	}

	return nil
}
