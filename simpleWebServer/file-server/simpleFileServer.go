package main

import "net/http"

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	// if the url is localhost:8080/ serve static files.
	http.Handle("/", http.FileServer(http.Dir("/Users/devraj/go")))
	// But if url is localhost:8080/ping , pong
	http.HandleFunc("/ping", handlePing)

	// It will start web server in port 8080.
	// You can use it in localhost:8080/
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

