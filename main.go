package main

import (
	"flag"
	"github.com/recep/add-string-to-filename/internal/cli"
	"github.com/recep/add-string-to-filename/internal/config"
	"github.com/recep/add-string-to-filename/internal/tools"
	"github.com/recep/add-string-to-filename/model"
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

	if opts.Undo {
		data, err := tools.Reader("./history/last.json")
		if err != nil {
			log.Fatalln(err)
		}

		file := model.File{}
		file.Undo(data)
		return
	}

	files, err := ioutil.ReadDir(opts.Path)
	if err != nil {
		log.Println(err)
	}

	// create a file object
	file := &model.File{}

	for _, f := range files {
		file.GetFileInfo(f.Name())

		// check options
		if opts.ShowFiles {
			cli.ShowFiles(*file)
			continue
		}

		if opts.Ext == file.ExtName || opts.Ext == "all" {
			if opts.AddEnd != "" {
				file.AddEnd(opts.AddEnd, opts.Path)
			}

			if opts.AddBeg != "" {
				file.AddBeginning(opts.AddEnd, opts.Path)
			}
		}

		if opts.Rename != "" && (opts.File == file.FullName || opts.File == file.Name) {
			file.Rename(opts.Rename, opts.Path)
		}
		// print file
		cli.PrintFile(*file)

		// reset file edited
		file.Edited = false
	}
}
