package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func countLetters(s io.Reader) (map[string]int, error) {
	buf := make([]byte, 4)
	out := map[string]int{}
	for {
		n, err := s.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func main() {
	s := "brown for over lazy dog"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(counts)
}
