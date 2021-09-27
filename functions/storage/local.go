package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Local struct {
	path string
	file string
	data string
}

func (f Local) Put() error {
	object, err := os.Create(f.file)
	if err != nil {
		log.Fatal(err)
	}

	defer object.Close()

	byteData := []byte(f.data)
	_, err = object.Write(byteData)
	fmt.Printf("%s created\n", f.file)
	fmt.Println("--------------------------------")
	return err
}

func (f Local) List() {
	dir := filepath.Dir(f.path)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("List of files at %s\n", f.path)
	fmt.Println("--------------------------------")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func (f Local) Get() ([]byte, error) {
	dat, err := os.ReadFile(f.file)
	return dat, err
}

func (f Local) Delete() error {
	err := os.Remove(f.file) // remove a single file
	fmt.Printf("%s deleted\n", f.file)
	fmt.Println("--------------------------------")
	return err
}

func newLocal(path, file, data string) Storage {
	return &Local{
		path: path,
		file: file,
		data: data,
	}
}
