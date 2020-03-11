package main

import (
	"fmt"
	"net/http"
)

type hotdog string

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("This-Key", "This is a test entry")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var dawg hotdog
	http.ListenAndServe(":8000", dawg)
}
