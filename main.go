package main

import (
	"flag"
	"github.com/recep/add-string-to-filename/internal/config"
	"github.com/recep/add-string-to-filename/internal/model"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fs := flag.NewFlagSet(filepath.Base(os.Args[0]), flag.ExitOnError)

	opts, err := config.ConfigureOptions(fs, os.Args[1:])
	if err != nil {
		config.PrintUsageErrorAndDie(err)
	}

	if opts.ShowHelp {
		config.PrintHelpAndDie()
	}

	files, err := ioutil.ReadDir(opts.Path)
	if err != nil {
		log.Println(err)
	}

	// create a file object
	file := &model.File{}

	for _, f := range files {
		file.GetFileInfo(f.Name())

		if opts.Ext == file.ExtName || opts.Ext == "all" {
			if opts.AddBeg != "" {
				err := file.AddEnd(opts.AddEnd, opts.Path)
				if err != nil {
					log.Fatalln(err)
				}
				return
			}

			if opts.AddBeg != "" {
				err := file.AddBeginning(opts.AddEnd, opts.Path)
				if err != nil {
					log.Fatalln(err)
				}
				return
			}
		}

		if opts.Rename != "" && (opts.File == file.FullName || opts.File == file.Name) {
			err := file.Rename(opts.Rename, opts.Path)
			if err != nil {
				log.Fatalln(err)
			}
			return
		}

	}
}
