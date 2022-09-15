package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func HandleRoot(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("HandleRoot!"))
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", HandleRoot)

	fmt.Printf("Starting server at port 8080\n")
	fmt.Println("working on port: ",port)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
