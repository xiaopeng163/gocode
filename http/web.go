package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	io.WriteString(w, "Hello world, this is version 1")
}
