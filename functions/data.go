package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Local struct {
	path string
	data string
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func (f *Local) SetPath(path string) error {
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

func (f *Local) SetData(data string) {
	f.data = data
}

func (f Local) Path() string {
	return f.path
}

func (f Local) Data() string {
	return f.data
}

func (f Local) Put(path string, data string) {
	file, err := os.Create(path)
	checkErr(err)

	defer file.Close()

	byteData := []byte(data)
	_, err = file.Write(byteData)
	checkErr(err)
	fmt.Println("File created")
}

func (f Local) List(path string) {
	dir := filepath.Dir(path)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func (f Local) Get(path string) string {
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(dat)
}

func (f Local) Delete(path string) error {
	err := os.Remove(path) // remove a single file
	return err
}

type Storage interface {
	Put(path string, data string)
	List(path string)
	Get(path string) string
	Delete(path string) error
	SetPath(path string) error
	SetData(data string)
	Path() string
	Data() string
}

func main() {
//	files := Local{}
	var files Storage
	files = &Local{}
	err := files.SetPath("/tmp/golang-test-file.txt")
	checkErr(err)
	files.SetData("my random data to a file")

	fmt.Println("File path:", files.Path(), "File data:", files.Data())
	files.Put(files.Path(), files.Data())
	files.List(files.Path())
	fmt.Println(files.Get(files.Path()))
	err = files.Delete(files.Path())
	checkErr(err)
}
