package main

import "fmt"

func getStorageKind(storage, path, file, data string) (Storage, error) {
	if storage == "linux" {
		return newLocal(path, file, data), nil
	}
        if storage == "cloud" {
                return newCloud(path, file, data), nil
        }
	return nil, fmt.Errorf("Wrong storage type passed")
}
