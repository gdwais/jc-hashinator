package main

import (
	"fmt"
	"net/http"
)

func logMessage(message string) {
	fmt.Println(message)
}

func logError(w http.ResponseWriter, message string) {
	fmt.Fprintf(w, "invalid http method")
}
