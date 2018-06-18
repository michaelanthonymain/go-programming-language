package main

import (
	"log"
	"net/http"

	"github.com/michaelanthonymain/go-programming-language/ch1/lissajous"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous.Lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
