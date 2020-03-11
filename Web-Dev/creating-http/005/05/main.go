package main

import (
	"io"
	"net/http"
)

func hotdog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hotdog")
}

func burger(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "burger")
}

func main() {

	http.Handle("/hotdog", http.HandlerFunc(hotdog))
	http.Handle("/burger", http.HandlerFunc(burger))

	http.ListenAndServe(":8000", nil)
}
