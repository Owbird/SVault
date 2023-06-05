package uifunctions

import (
	"github.com/owbird/svault/backend/ui_functions/dir"
	"github.com/owbird/svault/backend/ui_functions/engine"
)

type UIFunctions struct {
	dir.DirFunctions
	engine.EngineFunctions
}

func NewUIFunctions() *UIFunctions {
	return &UIFunctions{
		DirFunctions:    dir.DirFunctions{},
		EngineFunctions: engine.EngineFunctions{},
	}
}
