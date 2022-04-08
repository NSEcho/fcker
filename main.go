package main

import (
	"io"
	"net/http"
	"os"
)

const (
	thisPersonURL = "https://thispersondoesnotexist.com/image"
)

func main() {
	resp, err := http.Get(thisPersonURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	f, err := os.Create("slicica.jpeg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := io.Copy(f, resp.Body); err != nil {
		panic(err)
	}
}
