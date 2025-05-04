package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello, E4Devs!</h1>"))
	})

	http.ListenAndServe(":8080", nil)
}
// This is a simple HTTP server that responds with "Hello, E4Devs!" when accessed.