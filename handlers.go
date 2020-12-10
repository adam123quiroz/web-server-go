package main

import (
	"fmt"
	"net/http"
)

func HandlerRoot(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World")
}

func HandleHome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "this is the API Endpoint")
}
