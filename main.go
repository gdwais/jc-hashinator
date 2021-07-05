package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var port string = "8080"
var shuttingDown bool = false

func HashHandler(w http.ResponseWriter, r *http.Request) {
	if shuttingDown {
		httpOutput(w, "shutting down...")
		return
	}

	if r.Method == http.MethodGet {
		id := strings.TrimPrefix(r.URL.Path, "/hash/")
		i, _ := strconv.Atoi(id)
		hash := GetHash(i)
		httpOutput(w, hash)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		password := r.FormValue("password")
		i := AddRecord()
		ProcessAsyncRequest(i, password)
		httpOutput(w, strconv.Itoa(i))
		return
	} else {
		httpError(w)
		return
	}
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	if shuttingDown {
		httpOutput(w, "shutting down...")
		return
	}

	if r.Method == http.MethodGet {
		stats := GetStats()
		json.NewEncoder(w).Encode(stats)
	}
}

//TODO: rework this to use channel
func ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpError(w)
		return
	}

	shuttingDown = true

	if IsPending() {
		go func() {
			logMessage("waiting for processing to complete....")
			time.Sleep(5 * time.Second)
			logMessage("exiting")
			os.Exit(0)
		}()
	} else {
		logMessage("nothing is processing")
		logMessage("exiting")
		os.Exit(0)
	}
}

func runServer() {

	http.HandleFunc("/hash", HashHandler)
	http.HandleFunc("/hash/", HashHandler)
	http.HandleFunc("/stats", StatsHandler)

	http.HandleFunc("/shutdown", ShutdownHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	logMessage("starting server on port " + port)
	Store = make(map[int]Record)
	runServer()
}
