package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Message smth
type Message struct {
	Text string
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about/", about)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %s", r.URL.Path[1:])
}

func about(w http.ResponseWriter, r *http.Request) {

	m := Message{"Welcome JSON"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Write(b)
}
