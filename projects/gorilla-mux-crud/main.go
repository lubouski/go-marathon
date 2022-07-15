package main

import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler)

    fmt.Println("starting server on port 8080")
    log.Fatal(http.ListenAndServe(":8080",r))
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}
