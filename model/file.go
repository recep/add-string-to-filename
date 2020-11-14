package model

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	FullName string
	Name     string
	ExtName  string
	Path     string
	Edited   bool
}

func (f *File) String() string {
	return fmt.Sprintf("File: %s", f.FullName)
}

func (f *File) GetFileInfo(fileName string) {
	f.FullName = fileName
	f.ExtName = filepath.Ext(f.FullName)
	f.Name = strings.TrimSuffix(f.FullName, f.ExtName)
}

// Add() method adds string to the end of the file name.
func (f *File) Add(text, path string) error {
	oldPath := path + f.FullName

	f.Name += text
	f.Edited = true
	f.FullName = f.Name + f.ExtName

	newPath := path + f.FullName

	if err := os.Rename(oldPath, newPath); err != nil {
		return err
	}
	return nil
}
