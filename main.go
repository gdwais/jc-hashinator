package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var port string = "8080"

func PostHashHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		logError(w, "invalid_http_method")
		return
	}
	logMessage("Endpoint Hit: /hash")
}

func HashHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := strings.TrimPrefix(r.URL.Path, "/hash/")
		logMessage("Endpoint Hit GET: /hash/" + id)

	} else if r.Method == http.MethodPost {
		logMessage("Endpoint Hit POST: /hash")

	}
	logMessage("Endpoing Hit: /hash/1")

}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		stats := GetStats()
		json.NewEncoder(w).Encode(stats)
	}
}

func runServer() {
	http.HandleFunc("/hash/", HashHandler)
	http.HandleFunc("/stats", StatsHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	logMessage("starting server on port " + port)
	Store = make(map[int]Record)
	i1 := AddRecord()
	i2 := AddRecord()
	time.Sleep(10 * time.Second)
	CompleteRecord(i1, "Record1hash")
	CompleteRecord(i2, "Record2Hash")
	for key, value := range Store {
		fmt.Printf("%v value is %v\n", key, value)
	}
	runServer()
}
