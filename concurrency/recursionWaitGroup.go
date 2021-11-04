package main

import (
	"fmt"
	"sync"
	"time"
)

var g sync.WaitGroup

//func recur(iter int, g *sync.WaitGroup) {
func recur(iter int) {
    g.Add(1)
    defer g.Done()

    if iter <= 0 {
        return
    }

    go recur(iter-1)

    time.Sleep(time.Second * 10)
    fmt.Println(iter)

}

func main() {
    runs := 5

    recur(runs)

    g.Wait()
}
