package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	text := queryValues.Get("text")

	if len(text) > 0 {
		fmt.Fprintf(w, "%s\n", text)
		log.Println("text=" + text)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
