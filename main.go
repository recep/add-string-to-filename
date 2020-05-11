package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
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

		if *extName == "all" {
			ok = changer(path, exName, textName, name, extensionName)
		}

		if *extName == extensionName {
			ok = changer(path, exName, textName, name, extensionName)
		}
	}

	if !ok {
		fmt.Println("You do not have this file type!")
		return
	}

	printFiles(path, files)
}

func printFiles(path *string, files []os.FileInfo) {

	color.Blue("%-30s | %s\n", "Filename", "Changed")
	fmt.Println(strings.Repeat("-", 45))

	nFiles, err := ioutil.ReadDir(*path)
	if err != nil {
		log.Fatal(err)
	}

	for i, file := range nFiles {

		var c string

		if file.Name() != files[i].Name() {
			c = "+"
		}
		fmt.Printf("%-30s | %s\n", file.Name(), color.RedString(c))
		fmt.Println(strings.Repeat("·", 45))

	}
}

func changer(path, exName *string, textName []string, name, extensionName string) bool {
	newName := strings.Join(textName, ".")

	oldPath := *path + name
	newPath := *path + newName + *exName + "." + extensionName

	if err := os.Rename(oldPath, newPath); err != nil {
		log.Fatal(err)
	}
	return true
}
