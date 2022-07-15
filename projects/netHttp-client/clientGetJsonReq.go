package main

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
	"context"
)

func main() {
	// default client have no timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	
	// if you are making a PUT, POST, or PATCH request, specify the body of the request with the last parameter as an io.Reader
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		panic(err)
	}

	// you cat set any headers via the Headers field of the instance
	req.Header.Add("X-My-Client", "Learning Go")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: go %v", res.Status))
	}
	fmt.Println(res.Header.Get("Content-Type"))
	var data struct {
		UserID    int    `json:"userId"`
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)
}
