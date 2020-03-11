package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

type customWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Println(err)
	}

	io.Copy(os.Stdout, resp.Body)
}

func (customWriter) Write(bs []byte) (int, err) {
	return 1, nil
}
