package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		b, err := httputil.DumpRequest(request, true)
		if err == nil {
			fmt.Println(string(b))
		}

		writer.Write([]byte("ok"))
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
