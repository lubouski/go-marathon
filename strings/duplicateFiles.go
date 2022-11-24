package main

import (
	"fmt"
	"os"
	"log"
	"hash/crc32"
)

func main() {
	var fileSize os.FileInfo
	crc32q := crc32.MakeTable(0xD5828281)

	files, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	comparison := map[int64]int{}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileSize, err = file.Info()
		if err != nil {
			log.Fatal(err)
		}
		_, ok := comparison[fileSize.Size()]
		if !ok {
			comparison[fileSize.Size()] = 1
			continue
		}
		data, err := os.ReadFile(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%08x\n", crc32.Checksum(data, crc32q))		
	} 
}
