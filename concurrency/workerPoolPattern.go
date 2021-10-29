package main

import (
	"fmt"
)

func main() {
        jobs := make(chan int, 50)
        results := make(chan int, 50)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 50; i++ {
		jobs <- i
	}
	close(jobs)

	for  j := 0; j < 50; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n - 1) + fib(n - 2)
}
