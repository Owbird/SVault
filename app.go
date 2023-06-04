package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)

type App struct {
	ctx context.Context
}

type DirList struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
	Path  string `json:"path"`
	Type  string `json:"type"`
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetUserHome() string {
	home_dir, err := os.UserHomeDir()

	if err != nil {
		log.Println(err)
	}

	return home_dir
}

func (a *App) OpenFile(path string) {

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

func (a *App) GetDirs(p string) []DirList {
	var dir string

	dirList := []DirList{}

	if p == "/" {
		home_dir, err := os.UserHomeDir()

		if err != nil {
			log.Println(err)
		}

		dir = home_dir

	} else {
		dir = p
	}

	files, err := os.ReadDir(dir)

	if err != nil {
		log.Println(err)
	}

	for _, file := range files {

		path := path.Join(dir, file.Name())

		file_type := filepath.Ext(path)

		dirList = append(dirList, DirList{
			Name:  file.Name(),
			IsDir: file.IsDir(),
			Type:  file_type,
			Path:  path,
		})
	}

	return dirList

}
