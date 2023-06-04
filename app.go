package main

import (
	"context"

	uifunctions "github.com/owbird/svault/backend/ui_functions"
)

type App struct {
	Ctx         context.Context
	UIFunctions *uifunctions.UIFunctions
}

func NewApp() *App {
	return &App{
		UIFunctions: uifunctions.NewUIFunctions(),
	}
}

func (a *App) startup(ctx context.Context) {
	a.Ctx = ctx
}
