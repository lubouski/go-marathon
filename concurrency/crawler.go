package main

import (
	"fmt"
	"path/filepath"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	basedir := "/Users/Aliaksandr_Lubouski/path/tmp"

	jobs := make(chan map[string]int, 500)
	crawler(basedir, jobs)

	wg.Wait()
	close(jobs)

	result := make(map[string]int)
	for i := range jobs {
		for k,v := range i {
			result[k] = result[k] + v
		}
	}

	fmt.Println(result)

}

func crawler(basedir string, jobs chan map[string]int) {
	extensions := make(map[string]int)

	files, err := ioutil.ReadDir(basedir)
	if err != nil {
		log.Fatal(err)
	}


	for _, file := range files {
		if file.IsDir() == true {
			wg.Add(1)
			go func () {
				crawler(basedir + "/" + file.Name(),jobs)
				wg.Done()
			}()
		}

		extensions[filepath.Ext(file.Name())] = extensions[filepath.Ext(file.Name())] + 1
		time.Sleep(time.Millisecond * 50)
		fmt.Println(basedir + "/" + file.Name())
	}

	jobs <- extensions
}
