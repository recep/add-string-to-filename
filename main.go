package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//const exName = "-800x600."
func main() {

	dir := flag.String("dir", "", ".")
	exName := flag.String("s", "", "nil")

	//TODO special extension support
	//extName := flag.String("ext", "", "")

	//TODO error handling for empty command line argument

	flag.Parse()

	files, err := ioutil.ReadDir(*dir)

	if err != nil {
		panic(err)
	}

	for _, file := range files {

		name := file.Name()

		if name == "main.go" {
			continue
		}

		sName := strings.Split(name, ".")
		textName := sName[:len(sName)-1]
		extName := sName[len(sName)-1]

		var newName string

		for _, n := range textName {
			newName += n
		}

		newName += *exName + "." + extName

		err := os.Rename(file.Name(), newName)

		if err != nil {
			panic(err)
		}

		fmt.Println(name)
	}

}
