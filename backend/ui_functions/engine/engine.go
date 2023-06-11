package engine

import (
	"log"
	"os/exec"
)

type EngineFunctions struct {
}

func (ef *EngineFunctions) Encrypt(path string, pwd string) string {

	res, err := exec.Command("./build/bin/ee1", path, pwd).CombinedOutput()

	if err != nil {
		log.Println("Encrypt err ==> ", err)
	}

	return string(res)

}

func (ef *EngineFunctions) Decrypt() {}
