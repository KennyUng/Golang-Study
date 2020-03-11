package main

import (
	"io"
	"net/http"
)

/*
func hotdog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hotdog")
}

func burger(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "burger")
}
*/

func main() {

	http.HandleFunc("/hotdog", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "hotdog") })
	http.HandleFunc("/burger", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "burger") })

	http.ListenAndServe(":8000", nil)
}
