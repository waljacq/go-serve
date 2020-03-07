package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8181", "http service address")

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello\n")
}

func main() {
	flag.Parse()
	fmt.Printf("Starting on address : %s", *addr)
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
