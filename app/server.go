package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// get a unique pod identifier without using an environment variable
		podID := r.Header.Get("X-Pod-ID")
		w.Write([]byte("<h1>Hello, E4Devs! v7 - Deployment</h1><p>Pod ID: " + podID + "</p>"))
	})

	http.ListenAndServe(":8080", nil)
}
// This is a simple HTTP server that responds with "Hello, E4Devs! v7 - Deployment" when accessed.