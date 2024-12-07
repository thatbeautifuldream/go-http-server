package main

import (
	"flag"
	"fmt"
	"net/http"
)

var addr = flag.String("addr", ":3000", "http service address")

func main() {
	flag.Parse()
	http.HandleFunc("/", home)
	http.ListenAndServe(*addr, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, world")
}