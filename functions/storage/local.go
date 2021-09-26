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
	file string
	data string
}

func (f *Local) SetPath(path string) error {
	for _, l := range path {
		switch {
		case string(l) == "*":
			return errors.New("invalid linux path provided with symbol *")
		case string(l) == "?":
			return errors.New("invalid linux pfth provided with symbol ?")
		case string(l) == "\"":
			return errors.New("invalid linux path provided with symbol \"")
		}
	}
	f.path = path
	return nil
}

func (f *Local) SetFile(file string) error {
        for _, l := range file {
                switch {
                case string(l) == "*":
                        return errors.New("invalid linux filename provided with symbol *")
                case string(l) == "?":
                        return errors.New("invalid linux filename provided with symbol ?")
                case string(l) == "\"":
                        return errors.New("invalid linux filename provided with symbol \"")
                case string(l) == "/":
                        return errors.New("invalid linux filename provided with symbol /")
                }
        }
        f.file = file
        return nil
}

func (f *Local) SetData(data string) {
	f.data = data
}

func (f Local) GetPath() string {
	return f.path
}

func (f Local) GetFile() string {
        return f.file
}

func (f Local) GetData() string {
	return f.data
}

func (f Local) Put(path, file, data string) error {
	object, err := os.Create(path + file)
	if err != nil {
		log.Fatal(err)
	}

	defer object.Close()

	byteData := []byte(data)
	_, err = object.Write(byteData)
	fmt.Printf("%s created\n", path + file)
	fmt.Println("--------------------------------")
	return err
}

func (f Local) List(path string) {
	dir := filepath.Dir(path)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("List of files at %s\n", path)
	fmt.Println("--------------------------------")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func (f Local) Get(path, file string) ([]byte, error) {
	dat, err := os.ReadFile(path + file)
	return dat, err
}

func (f Local) Delete(path, file string) error {
	err := os.Remove(path + file) // remove a single file
	fmt.Printf("%s deleted at dir %s\n", file, path)
	fmt.Println("--------------------------------")
	return err
}

func newLocal() Storage {
	return &Local{
		path: "/tmp/",
		file: "golang-test-filename.txt",
		data: "my random data to a file",
	}
}
