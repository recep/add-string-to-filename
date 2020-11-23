package config

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

var usageStr = `
Usage: editor-go
Options:
	-p, --path <path your folder> 
	-ae <string>
	-ab <string>
	-f, --file <file name>
	-r, --rename <new file name>
	-u, --undo 
Common Options: 
	-help
`

func PrintUsageErrorAndDie(err error) {
	log.Println(err)
	fmt.Println(usageStr)
	os.Exit(1)
}

func PrintHelpAndDie() {
	fmt.Println(usageStr)
	os.Exit(0)
}

type Options struct {
	Path     string `json:"path"`
	AddEnd   string `json:"add_end"`
	AddBeg   string `json:"add_beg"`
	Ext      string `json:"ext"`
	File     string `json:"file"`
	Rename   string `json:"rename"`
	Undo     bool   `json:"undo"`
	ShowHelp bool   `json:"show_help"`
}

func ConfigureOptions(fs *flag.FlagSet, args []string) (*Options, error) {
	opts := &Options{}

	fs.StringVar(&opts.Path, "p", "", "Folder path")
	fs.StringVar(&opts.Path, "path", "", "Folder Path")
	fs.StringVar(&opts.AddEnd, "ae", "", "Add string to the end of the file name")
	fs.StringVar(&opts.AddBeg, "ab", "", "Add string to the beginning of the file name")
	fs.StringVar(&opts.Ext,"ext","all","Extension name of file")
	fs.StringVar(&opts.File, "f", "", "specific file")
	fs.StringVar(&opts.File, "file", "", "specific file")
	fs.StringVar(&opts.Rename, "r", "", "new file name")
	fs.StringVar(&opts.Rename, "rename", "", "new file name")
	fs.BoolVar(&opts.Undo, "u", false, "undo last change")
	fs.BoolVar(&opts.Undo, "undo", false, "undo last change")
	fs.BoolVar(&opts.ShowHelp, "h", false, "show help message")
	fs.BoolVar(&opts.ShowHelp, "help", false, "show help message")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if !opts.ShowHelp && opts.Path == "" {
		err := errors.New("please specify all arguments")
		return nil, err
	}

	return opts, nil
}
