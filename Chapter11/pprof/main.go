package main

import (
	"net/http"
	_ "net/http/pprof"
)

//https://golang.org/pkg/net/http/pprof/
func main() {
	http.HandleFunc("/", func(wr http.ResponseWriter, req *http.Request) {
		wr.Write([]byte("hello,world!!"))
	})
	http.ListenAndServe(":8888", nil)

}
