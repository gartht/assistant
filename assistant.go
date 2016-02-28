package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe("0.0.0.0:3000", nil)
}

func indexHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Hello World")
}
