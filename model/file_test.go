package model

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()

	exitVal := m.Run()
	if exitVal == 0 {
		teardown()
	}
	os.Exit(exitVal)
}

func TestFile_AddEnd(t *testing.T) {
	files, err := ioutil.ReadDir("./folder")
	if err != nil {
		log.Fatalln(err)
	}

	var file File
	text := "go"

	for _, f := range files {
		file.GetFileInfo(f.Name())
		want := file.Name + text + file.ExtName
		got := file.AddEnd(text, "folder/")
		assert.Equal(t, want, got, "they should be equal!")
	}
}

func TestFile_AddBeginning(t *testing.T) {
	files, err := ioutil.ReadDir("./folder")
	if err != nil {
		log.Fatalln(err)
	}

	var file File
	text := "go"

	for _, f := range files {
		file.GetFileInfo(f.Name())
		want := text + file.Name + file.ExtName
		got := file.AddBeginning(text, "folder/")
		assert.Equal(t, want, got, "they should be equal!")
	}
}

func TestFile_Rename(t *testing.T) {
	files, err := ioutil.ReadDir("./folder")
	if err != nil {
		log.Fatalln(err)
	}

	var file File
	text := "newfilename"

	// Test: changing only the first file's name
	for _, f := range files {
		file.GetFileInfo(f.Name())
		want := text + file.ExtName
		got := file.Rename(text, "folder/")
		assert.Equal(t, want, got, "they should be equal!")
		break
	}
}

func setup() {
	err := os.Mkdir("./folder", 0755)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.Mkdir("./history", 0755)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = os.Create("./history/last.json")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = os.Create("./folder/example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = os.Create("./folder/example2.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func teardown() {
	err := os.RemoveAll("./folder")
	if err != nil {
		log.Fatalln(err)
	}

	err = os.RemoveAll("./history")
	if err != nil {
		log.Fatalln(err)
	}
}
