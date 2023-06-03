package main

import (
	"context"
	"log"
	"os"
	"path"
)

type App struct {
	ctx context.Context
}

type DirList struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
	Path  string `json:"path"`
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

		dirList = append(dirList, DirList{
			Name:  file.Name(),
			IsDir: file.IsDir(),
			Path:  path.Join(dir, file.Name()),
		})
	}

	return dirList

}
