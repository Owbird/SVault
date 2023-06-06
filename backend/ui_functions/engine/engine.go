package engine

import (
	"os/exec"
)

type EngineFunctions struct {
}

func (ef *EngineFunctions) Encrypt(path string) {

	exec.Command("./build/bin/ee1", path, "TEST_KEY").CombinedOutput()
}

func (ef *EngineFunctions) Decrypt() {}
