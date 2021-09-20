package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"os"
)


func main() {

	fmt.Println("Pleae enter five parameters according to README:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	param := strings.Split(input, " ")

	storage, _ := getStorageKind(param[0])

        setStorage(storage, param[1], param[2], param[3])

	switch {
	case param[4] == "put":
		putStorage(storage)
	case param[4] == "get":
		getStorage(storage)
	case param[4] == "list":
		listStorage(storage)
	case param[4] == "delete":
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
