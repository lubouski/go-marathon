package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"io/ioutil"
)

type File struct {
	path string
	data string
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func (f *File) SetPath(path string) error {
	for _, l := range path {
		switch {
		case string(l) == "*":
			return errors.New("invalid linux filename provided with symbol *")
		case string(l) == "?":
			return errors.New("invalid linux filename provided with symbol ?")
		case string(l) == "\"":
			return errors.New("invalid linux filename provided with symbol \"")
		}
	}
	f.path = path
	return nil
}

func (f *File) SetData(data string) {
	f.data = data
}

func (f *File) Path() string {
	return f.path
}

func (f *File) Data() string {
	return f.data
}

func (f *File) Put(path string, data string) {
	file, err := os.Create(path)
	checkErr(err)

	defer file.Close()

	byteData := []byte(data)
	_, err = file.Write(byteData)
	checkErr(err)
	fmt.Println("File created")
}

func (f *File) List(path string) {
	dir := filepath.Dir(path)
	files, err := ioutil.ReadDir(dir)
	checkErr(err)
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func (f *File) Get(path string) {
	dat, err := os.ReadFile(path)
	checkErr(err)
	fmt.Print(string(dat))
}

func main() {
	files := File{}
	err := files.SetPath("/tmp/golang-test-file.txt")
	checkErr(err)
	files.SetData("my random data to a file")
	fmt.Println("File path:", files.Path(), "File data:", files.Data())
	files.Put(files.Path(), files.Data())
	files.List(files.Path())
	files.Get(files.Path())
}
