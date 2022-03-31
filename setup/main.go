package main

import (
	"fmt"
	"os"
)

const perms = 0777

func msgDir() string {
	ret, _ := os.UserHomeDir()
	return ret + "/gomsg/msg/"
}

func usrDir() string {
	ret, _ := os.UserHomeDir()
	return ret + "/gomsg/usr/"
}

func createDir(path string) {
	err := os.MkdirAll(path, perms)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	for _, dir := range []string{msgDir(), usrDir()} {
		createDir(dir)
	}
}
