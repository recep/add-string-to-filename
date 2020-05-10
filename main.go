package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

	printFiles(path, files)
}

func printFiles(path *string, files []os.FileInfo) {

	fmt.Printf("%-30s | %s\n", "Filename", "Changed")
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

		fmt.Printf("%-30s | %s\n", file.Name(), c)
		fmt.Println(strings.Repeat("·", 45))

	}
}
