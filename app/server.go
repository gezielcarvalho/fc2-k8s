package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("<h1>Hello, E4Devs! v12 - Deployment</h1><p>Served by: %s</p>", hostname)
		w.Write([]byte(message))
	})

	http.ListenAndServe(":8000", nil)
}
