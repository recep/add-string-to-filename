package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	path := flag.String("path", "", "PATH")
	exName := flag.String("s", "", "text")
	extName := flag.String("ext", "", "extension name of file")

	flag.Parse()

	if *path == "" || *exName == "" || *extName == "" {
		fmt.Println(`please enter all commands 
		go run main.go -path=tmp -s=string -ext=txt`)
		return
	}

	files, err := ioutil.ReadDir(*path)
	if err != nil {
		panic(err)
	}

	var ok bool

	for _, file := range files {

		name := file.Name()

		if name == "main.go" {
			continue
		}

		sName := strings.Split(name, ".")
		textName := sName[:len(sName)-1]
		extensionName := sName[len(sName)-1]

		var (
			newName string
			oldPath string
			newPath string
		)

		if *extName == "all" {

			//Birden fazla nokta içeren dosya varsa ayrılan isimlerini birleştirme
			newName = strings.Join(textName, ".")

			oldPath = *path + name
			newPath = *path + newName + *exName + "." + extensionName

			if err := os.Rename(oldPath, newPath); err != nil {
				panic(err)
			}
			ok = true
		}

		if *extName == extensionName {

			newName = strings.Join(textName, ".")

			oldPath = *path + name
			newPath = *path + newName + *exName + "." + extensionName

			if err := os.Rename(oldPath, newPath); err != nil {
				panic(err)
			}
			ok = true
		}
	}

	if !ok {
		fmt.Println("You do not have this file type!")
		return
	}

	nFiles, err := ioutil.ReadDir(*path)
	if err != nil {
		panic(err)
	}

	fmt.Println("Your new files:\n", strings.Repeat("-", 25))
	for _, file := range nFiles {
		fmt.Println(file.Name())
	}

}
