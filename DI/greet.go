package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writter io.Writer, name string) {
	fmt.Fprintf(writter, "Hello, %s\n", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Omar")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
