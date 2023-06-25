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

func (ef *EngineFunctions) Decrypt(path string, pwd string) string {
	res, err := exec.Command("./build/bin/ee2", path, pwd).CombinedOutput()

	log.Println(path, pwd)

	if err != nil {
		log.Println("Decrypt err ==> ", err)
	}

	return string(res)
}
