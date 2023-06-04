package uifunctions

import (
	"github.com/owbird/svault/backend/ui_functions/dir"
)

type UIFunctions struct {
	dir.DirFunctions
}

func NewUIFunctions() *UIFunctions {
	return &UIFunctions{
		DirFunctions: dir.DirFunctions{},
	}
}
