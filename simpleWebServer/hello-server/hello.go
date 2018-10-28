package main

import (
	"net/http"
	"strings"
)
// hit url: localhost:8080/devraj , in browser we will see Hello devraj
func printHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message

	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", printHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

