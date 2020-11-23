package cli

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/recep/add-string-to-filename/internal/model"
	"strings"
)

// Print
func PrintFile(file model.File) {
	if file.Edited {
		fmt.Printf("%-25s | %-25s | %s\n", "NAME", "OLD NAME", "STATUS")
		fmt.Println(strings.Repeat("-", 50))
		e := color.GreenString("EDITED")
		fmt.Printf("%-25s | %-25s | %s\n", file.FullName, file.OldName, e)
		return
	}
}

// Show all files
func ShowFiles(file model.File) {
	fmt.Println(strings.Repeat("-", 65))
	fmt.Printf("%-25s | %-25s | %s\n", file.FullName, file.Name, file.ExtName)
}
