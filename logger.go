package main

import (
	"fmt"
	"net/http"
)

func LogMessage(message string) {
	fmt.Println(message)
}

func LogError(message string) {
	fmt.Println("ERROR! : " + message)
}

func HttpError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	HttpOutput(w, "invalid http method")
}

func HttpOutput(w http.ResponseWriter, message string) {
	fmt.Fprintf(w, message)
}
