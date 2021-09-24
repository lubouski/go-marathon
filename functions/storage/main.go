package main

import (
	"flag"
	"fmt"
	"log"
)


func main() {
	var provider string
	var path string
	var file string
	var data string
	var command string

	flag.StringVar(&provider, "provider", "cloud", "storage provider")
	flag.StringVar(&path, "path", "/tmp", "path or bucket")
	flag.StringVar(&file, "file", "test.txt", "file name")
	flag.StringVar(&data, "data", "secret", "data or file to upload to cloud")
	flag.StringVar(&command, "command", "get", "get,put, list or delete command")
	flag.Parse()

	storage, _ := getStorageKind(provider)

        setStorage(storage, path, file, data)

	switch {
	case command == "put":
		putStorage(storage)
	case command == "get":
		getStorage(storage)
	case command == "list":
		listStorage(storage)
	case command == "delete":
		deleteStorage(storage)
	}
}

func setStorage(s Storage, path string, file string, data string) {
	err := s.SetPath(path)
	if err != nil {
		log.Fatal(err)
	}
	err = s.SetFile(file)
	if err != nil {
		log.Fatal(err)
	}
	s.SetData(data)
}

func putStorage(s Storage) {
	err := s.Put(s.GetPath(), s.GetFile(), s.GetData())
	if err != nil {
		log.Fatal(err)
	}
}

func listStorage(s Storage) {
	s.List(s.GetPath())
}

func getStorage(s Storage) {
	data, err := s.Get(s.GetPath(), s.GetFile())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func deleteStorage(s Storage) {
	err := s.Delete(s.GetPath(), s.GetFile())
	if err != nil {
		log.Fatal(err)
	}
}
