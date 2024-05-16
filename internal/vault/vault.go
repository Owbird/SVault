package vault

import (
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type Vault struct{}

type Dir struct {
	Name    string `json:"name"`
	IsVault bool   `json:"isVault"`
	Path    string `json:"path"`
}

type VaultFunctions struct{}

func NewVaultFunctions() *VaultFunctions {
	return &VaultFunctions{}
}

func (vf *VaultFunctions) GetUserHome() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
	}

	return homeDir
}

func (vf *VaultFunctions) OpenFile(path string) {
	var cmd string
	var args []string

	if runtime.GOOS == "linux" {
		cmd = "xdg-open"
		args = []string{
			path,
		}
	} else if runtime.GOOS == "darwin" {
		cmd = "open"
		args = []string{
			path,
		}
	} else {
		cmd = "cmd"
		args = []string{
			"/c",
			"start",
			path,
		}
	}

	exec.Command(cmd, args...).Run()
}

func (vf *VaultFunctions) GetDirs(p string) []Dir {
	var dir string

	dirList := []Dir{}

	if p == "/" {
		dir = vf.GetUserHome()
	} else {
		dir = p
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Println(err)
	}

	for _, file := range files {

		path := path.Join(dir, file.Name())

		dirList = append(dirList, Dir{
			Name:    file.Name(),
			IsVault: file.IsDir(),
			Path:    path,
		})

	}

	return dirList
}

func (vf *VaultFunctions) MoveToVault(p string) error {
	userHome := vf.GetUserHome()

	vaultPath := strings.ReplaceAll(p, userHome, "")

	vaultPath = path.Join(".vault", vaultPath)

	err := os.MkdirAll(filepath.Dir(vaultPath), os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	return os.Rename(p, vaultPath)
}

func (vf *VaultFunctions) MoveFromVault(p string) error {
	userHome := vf.GetUserHome()

	normPath := strings.ReplaceAll(p, ".vault", userHome)

	return os.Rename(p, normPath)
}

func (vf *VaultFunctions) DeleteFile(p string) {
	err := os.Remove(filepath.Join(p))
	if err != nil {
		log.Fatalln(err)
	}
}
