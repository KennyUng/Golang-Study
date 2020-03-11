package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hotdog")
}

type burger int

func (b burger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "burger")
}

func main() {
	var h hotdog
	var b burger

	mux := http.NewServeMux()
	mux.Handle("/hotdog", h)
	mux.Handle("/burger", b)

	http.ListenAndServe(":8000", mux)
}
