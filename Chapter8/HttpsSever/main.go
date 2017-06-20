package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello,htts"))
	})
	http.ListenAndServeTLS(":8888", "server.crt", "server.key", nil)
}
