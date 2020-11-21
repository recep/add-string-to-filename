package main

import (
	"flag"
	"fmt"
	"github.com/recep/add-string-to-filename/internal/model"
	"io/ioutil"
	"log"
)

func main() {

	path := flag.String("path", "", "PATH")
	fName := flag.String("f","","file name")
	ae := flag.String("ae", "", "adds string to the end of the file name")
	ab := flag.String("ab", "", "adds string to the beginning of the file name.")
	rn := flag.String("rn","","rename file name")
	ext := flag.String("ext", "", "extension name of file")

	flag.Parse()

	if *path == ""  {
		fmt.Println(`please enter all commands 
		go run main.go 
		-path folder/ -ae string -ext .txt
		-path folder/ -ab string -ext .txt`)
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

		if *ext == file.ExtName || *ext == ".all" {
			err := file.AddEnd(*ae, *path)
			if err != nil {
				log.Fatalln(err)
			}

			err = file.AddBeginning(*ab, *path)
			if err != nil {
				log.Fatalln(err)
			}
		}

		if file.FullName == *fName && *rn != "" {
			err := file.Rename(*rn,*path,*file)
			if err != nil {
				log.Fatalln(err)
			}
		}
		fmt.Println(file)
	}
}
