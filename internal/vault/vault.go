package vault

import (
	"github.com/Owbird/SVault-Engine/pkg/models"
	"github.com/Owbird/SVault-Engine/pkg/vault"
)

type Vault struct{}

type Dir struct {
	Name    string `json:"name"`
	IsVault bool   `json:"isVault"`
	Path    string `json:"path"`
}

type VaultFunctions struct {
	vault *vault.Vault
}

func NewVaultFunctions() *VaultFunctions {
	return &VaultFunctions{
		vault: vault.NewVault(),
	}
}

func (vf *VaultFunctions) CreateVault(name, pwd string) error {
	err := vf.vault.Create(name, pwd)

	if err != nil {
		return err
	}

	return nil
}

func (vf *VaultFunctions) ListVaults() ([]models.Vault, error) {

	return vf.vault.List()

}
