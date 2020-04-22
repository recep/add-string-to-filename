package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	path := flag.String("path", "", "a string")
	exName := flag.String("s", "", "a string")
	extName := flag.String("ext", "", "a string")

	flag.Parse()

	if *path == "" || *exName == "" || *extName == "" {
		fmt.Println(`please enter all commands 
		go run main.go -path=temp -s=string -ext=txt`)
		return
	}

	folder, err := os.Open(*path)
	if err != nil {
		panic(err)
	}

	files, err := folder.Readdir(0)
	if err != nil {
		panic(err)
	}

	if err := folder.Close(); err != nil {
		panic(err)
	}

	for _, file := range files {

		name := file.Name()

		if name == "main.go" {
			continue
		}

		sName := strings.Split(name, ".")
		textName := sName[:len(sName)-1]
		extensionName := sName[len(sName)-1]

		if *extName == extensionName {

			var newName string

			for _, n := range textName {
				newName += n
			}

			oldPath := *path + "/" + name
			newPath := *path + "/" + newName + *exName + "." + extensionName

			err := os.Rename(oldPath, newPath)
			if err != nil {
				panic(err)
			}
		}

	}

}
