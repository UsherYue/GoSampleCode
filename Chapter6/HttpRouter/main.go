package main

import (
	"io"
	"log"
	"net/http"
)

func Router1(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, router1!\n")
}

func Router2(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, router2!\n")
}

func Router3(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, router3!\n")
}

func main() {
	http.HandleFunc("/router1", Router1)
	http.HandleFunc("/router2", Router2)
	http.HandleFunc("/router3", Router3)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
