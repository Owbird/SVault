package dir

import (
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type Dir struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
	Path  string `json:"path"`
}

type DirFunctions struct {
}

func (df *DirFunctions) GetUserHome() string {
	home_dir, err := os.UserHomeDir()

	if err != nil {
		log.Println(err)
	}

	return home_dir
}

func (df *DirFunctions) OpenFile(path string) {
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

func (df *DirFunctions) GetDirs(p string) []Dir {
	var dir string

	dirList := []Dir{}

	if p == "/" {

		dir = df.GetUserHome()

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
			Name:  file.Name(),
			IsDir: file.IsDir(),
			Path:  path,
		})

	}

	return dirList
}

func (df *DirFunctions) MoveToVault(p string) error {
	user_home := df.GetUserHome()

	vault_path := strings.ReplaceAll(p, user_home, "")

	vault_path = path.Join(".vault", vault_path)

	err := os.MkdirAll(filepath.Dir(vault_path), os.ModePerm)

	if err != nil {
		log.Println(err)
	}

	return os.Rename(p, vault_path)

}

func (df *DirFunctions) MoveFromVault(p string) error {
	user_home := df.GetUserHome()

	norm_path := strings.ReplaceAll(p, ".vault", user_home)

	return os.Rename(p, norm_path)

}

func (df *DirFunctions) DeleteFile(p string) {
	err := os.Remove(filepath.Join(p))

	if err != nil {
		log.Fatalln(err)
	}
}
