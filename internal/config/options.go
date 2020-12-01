package config

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

var usageStr = `
Usage: File Name Editor-GO
Options:
	-p, --path <path>			path your folder 
	-ae <string>				Add string to the end of the file name
	-ab <string>				Add string to the beginning of the file name
	-f, --file <file name>			specific file 
	-rn, --rename <new file name>		new file name
	-u, --undo 				undo last change
	-ls					show files
Common Options: 
	-help					show help
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
	Path      string `tools:"path"`
	AddEnd    string `tools:"add_end"`
	AddBeg    string `tools:"add_beg"`
	Ext       string `tools:"ext"`
	File      string `tools:"file"`
	Rename    string `tools:"rename"`
	Undo      bool   `tools:"undo"`
	ShowFiles bool   `tools:"show_files"`
	ShowHelp  bool   `tools:"show_help"`
}

func ConfigureOptions(fs *flag.FlagSet, args []string) (*Options, error) {
	opts := &Options{}

	fs.StringVar(&opts.Path, "p", "", "Folder path")
	fs.StringVar(&opts.Path, "path", "", "Folder Path")
	fs.StringVar(&opts.AddEnd, "ae", "", "Add string to the end of the file name")
	fs.StringVar(&opts.AddBeg, "ab", "", "Add string to the beginning of the file name")
	fs.StringVar(&opts.Ext, "ext", "all", "Extension name of file")
	fs.StringVar(&opts.File, "f", "", "specific file")
	fs.StringVar(&opts.File, "file", "", "specific file")
	fs.StringVar(&opts.Rename, "rn", "", "new file name")
	fs.StringVar(&opts.Rename, "rename", "", "new file name")
	fs.BoolVar(&opts.Undo, "u", false, "undo last change")
	fs.BoolVar(&opts.Undo, "undo", false, "undo last change")
	fs.BoolVar(&opts.ShowFiles, "ls", false, "show files")
	fs.BoolVar(&opts.ShowHelp, "h", false, "show help message")
	fs.BoolVar(&opts.ShowHelp, "help", false, "show help message")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if !opts.ShowHelp && opts.Path == "" && !opts.Undo  {
		err := errors.New("please specify all arguments")
		return nil, err
	}

	return opts, nil
}
