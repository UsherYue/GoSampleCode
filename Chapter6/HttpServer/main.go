package main

import (
	"io"
	"log"
	"net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.Handle("/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":9999", nil))
}
