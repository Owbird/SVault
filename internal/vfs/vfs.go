package vfs

import "github.com/Owbird/SVault-Engine/pkg/filesystem"

type VFSFunctions struct{}

func NewVFSFunctions() *VFSFunctions {
	return &VFSFunctions{}
}

func (vfsf *VFSFunctions) Mount() {
	go filesystem.Mount()
}
