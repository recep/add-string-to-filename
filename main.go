package main

import (
	"flag"
	"fmt"
	"github.com/recep/add-string-to-filename/model"
	"io/ioutil"
	"log"
)

func main() {

	path := flag.String("path", "", "PATH")
	text := flag.String("s", "", "text")
	ext := flag.String("ext", "", "extension name of file")

	flag.Parse()

	if *path == "" || *text == "" || *ext == "" {
		fmt.Println(`please enter all commands 
		go run main.go -path=folder/ -s=string -ext=.txt`)
		return
	}

	files, err := ioutil.ReadDir(*path)
	if err != nil {
		log.Println(err)
	}

	// create an empty file object
	file := &model.File{}

	for _, f := range files {
		file.GetFileInfo(f.Name())

		if *ext == file.ExtName {
			err := file.Add(*text, *path)
			if err != nil {
				log.Fatalln(err)
			}
		}

		if *ext == ".all" {
			err := file.Add(*text, *path)
			if err != nil {
				log.Fatalln(err)
			}
		}

		fmt.Println(file)
	}
}
