package main

import "fmt"

func getStorageKind(storage string) (Storage, error) {
	if storage == "linux" {
		return newLocal(), nil
	}
        if storage == "cloud" {
                return newCloud(), nil
        }
	return nil, fmt.Errorf("Wrong storage type passed")
}
