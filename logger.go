package main

import (
	"fmt"
	"net/http"
)

func logMessage(message string) {
	fmt.Println(message)
}

func logError(message string) {
	fmt.Println("ERROR! : " + message)
}

func httpError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	httpOutput(w, "invalid http method")
}

func httpOutput(w http.ResponseWriter, message string) {
	fmt.Fprintf(w, message)
}
