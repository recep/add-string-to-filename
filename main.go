package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const exName = "-800x600."

func main() {

	args := os.Args[1:]

	files, err := ioutil.ReadDir(args[0])

	if err != nil {
		panic(err)
	}

	for _, file := range files {

		name := file.Name()

		if name == "main.go" {
			continue
		}

		sName := strings.Split(name, ".")
		newName := sName[0] + exName + sName[1]

		err := os.Rename(file.Name(), newName)

		if err != nil {
			panic(err)
		}

		fmt.Println(name)
	}

}
