package model

import (
	"fmt"
	"github.com/recep/add-string-to-filename/internal/json"
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

func (f *File) String() string {
	return fmt.Sprintf("File: %s", f.FullName)
}

func (f *File) GetFileInfo(file string) {
	f.FullName = file
	f.ExtName = filepath.Ext(f.FullName)
	f.Name = strings.TrimSuffix(f.FullName, f.ExtName)
}

// AddEnd() method adds string to the end of the file name.
func (f *File) AddEnd(text, path string) error {
	oldPath := path + f.FullName
	f.OldName = f.FullName

	f.Name += text
	f.Edited = true
	f.FullName = f.Name + f.ExtName
	f.ModTime = time.Now().String()

	newPath := path + f.FullName

	if err := os.Rename(oldPath, newPath); err != nil {
		return err
	}

	err := json.Writer(f, "./history/last.json")
	if err != nil {
		return err
	}

	return nil
}

// AddBeginning() method adds string to beginning of the file name.
func (f *File) AddBeginning(text, path string) error {
	oldPath := path + f.FullName
	f.OldName = f.FullName

	text += f.Name
	f.Name = text
	f.Edited = true
	f.FullName = f.Name + f.ExtName
	f.ModTime = time.Now().String()

	newPath := path + f.FullName

	if err := os.Rename(oldPath, newPath); err != nil {
		return err
	}

	err := json.Writer(f, "./history/last.json")
	if err != nil {
		return err
	}

	return nil
}

func (f *File) Rename(text,path string, file File) error {
	oldPath := path + f.FullName
	f.OldName = f.FullName

	f.Name = text
	f.FullName = f.Name + f.ExtName
	f.Edited = true
	f.ModTime = time.Now().String()

	newPath := path + f.FullName

	if err := os.Rename(oldPath, newPath); err != nil {
		return err
	}

	return nil
}
