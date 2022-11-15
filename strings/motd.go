package main

import (
	"fmt"
	"os"
	"io"
	"strings"
	"log"
	"math/rand"
)

func bufferedReader(r io.Reader) (map[int]string, error) {
	buf := make([]byte, 2048)
	motd := map[int]string{}

	var strBuf strings.Builder
	var count int

	for {
		n, err := r.Read(buf)
		for _,ch := range buf[:n] {
			if ch == '\n' {
				fmt.Println("New line!")
				motd[count] = strBuf.String()
				strBuf.Reset()
				count++
				continue
			}
			strBuf.WriteByte(ch)	
		}
		if err == io.EOF {
			return motd, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func main() {
	file, err := os.Open("new-line.txt")
	if err != nil {
		log.Fatal(err)
	}
	motd, err := bufferedReader(file)
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(33)
	fmt.Println(motd[rand.Intn(len(motd))])
}
