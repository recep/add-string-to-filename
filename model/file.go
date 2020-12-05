package model

import (
	"encoding/json"
	"github.com/recep/add-string-to-filename/internal/tools"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type File struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	ExtName  string `json:"ext_name"`
	Path     string `json:"path"`
	Edited   bool   `json:"edited"`
	OldName  string `json:"old_name"`
	ModTime  string `json:"mod_time"`
}

//func (f *File) String() string {
//	return fmt.Sprintf("File: %s", f.FullName)
//}

// Get file info
func (f *File) GetFileInfo(file string) {
	f.FullName = file
	f.ExtName = filepath.Ext(f.FullName)
	f.Name = strings.TrimSuffix(f.FullName, f.ExtName)
}

// AddEnd() method adds string to the end of the file name.
func (f *File) AddEnd(text, path string) {
	oldPath := path + f.FullName
	f.OldName = f.FullName

	f.Name += text
	f.Edited = true
	f.FullName = f.Name + f.ExtName
	f.Path = path
	f.ModTime = time.Now().String()

	newPath := path + f.FullName

	if err := os.Rename(oldPath, newPath); err != nil {
		log.Fatalln(err)
	}

	err := tools.Writer(f, "./history/last.json")
	if err != nil {
		log.Fatalln(err)
	}
}

// AddBeginning() method adds string to beginning of the file name.
func (f *File) AddBeginning(text, path string) {
	oldPath := path + f.FullName
	f.OldName = f.FullName

	text += f.Name
	f.Name = text
	f.Edited = true
	f.Path = path
	f.FullName = f.Name + f.ExtName
	f.ModTime = time.Now().String()

	newPath := path + f.FullName

	if err := os.Rename(oldPath, newPath); err != nil {
		log.Fatalln(err)
	}

	err := tools.Writer(f, "./history/last.json")
	if err != nil {
		log.Fatalln(err)
	}
}

// Rename file
func (f *File) Rename(text, path string) {
	oldPath := path + f.FullName
	f.OldName = f.FullName

	f.Name = text
	f.FullName = f.Name + f.ExtName
	f.Edited = true
	f.Path = path
	f.ModTime = time.Now().String()

	newPath := path + f.FullName

	if err := os.Rename(oldPath, newPath); err != nil {
		log.Fatalln(err)
	}

	if err := tools.Writer(f, "./history/last.json"); err != nil {
		log.Fatalln(err)
	}
}

func (f *File) Undo(data []byte) {

	if err := json.Unmarshal(data, f); err != nil {
		log.Fatalln(err)
	}

	oldPath := f.Path + f.FullName

	f.Name = strings.TrimSuffix(f.OldName, f.ExtName)
	f.FullName, f.OldName = f.OldName, f.FullName

	newPath := f.Path + f.FullName

	if err := os.Rename(oldPath, newPath); err != nil {
		log.Fatalln(err)
	}

	if err := tools.Writer(f, "./history/last.json"); err != nil {
		log.Fatalln(err)
	}
}
