package main

import (
	"log"
	"fmt"
)

func main() {
	linux, _ := getStorageKind("linux")
	cloud, _ := getStorageKind("cloud")

	setStorage(linux, "/tmp/", "my-test-golang.txt", "secret data for cloud")
	putStorage(linux)

	setStorage(cloud, "lubouski-golang", "golang-test.txt", "/tmp/my-test-golang.txt")
	putStorage(cloud)

	listStorage(linux)
	listStorage(cloud)

	getStorage(linux)
	getStorage(cloud)

	deleteStorage(linux)
	deleteStorage(cloud)
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
