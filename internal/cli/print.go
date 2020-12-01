package cli

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/recep/add-string-to-filename/model"
	"strings"
)

// Print
func PrintFile(file model.File) {
	if file.Edited {
		x := color.RedString("*")
		fmt.Printf("%s %-25s | %s\n", file.FullName, x, file.OldName)
		fmt.Println(strings.Repeat("-", 65))
	}
}

// Show all files
func ShowFiles(file model.File) {
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("%-30s | %-25s \n", file.FullName, file.ExtName)
}
