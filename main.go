package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func HashHandler(w http.ResponseWriter, r *http.Request) {
	if shuttingDown {
		HttpOutput(w, "shutting down...")
		return
	}

	if r.Method == http.MethodGet {
		id := strings.TrimPrefix(r.URL.Path, "/hash/")
		i, _ := strconv.Atoi(id)
		hash := GetHash(i)
		HttpOutput(w, hash)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		password := r.FormValue("password")
		i := AddRecord()
		go func() {
			encryptJobs <- IdAndValue{Id: i, Value: password}
			encryptResults := <-encryptResults
			completeRecordJobs <- encryptResults
			complete := <-completeRecordResults
			if complete {
				LogMessage(strconv.Itoa(i) + " completed")
			}
		}()
		HttpOutput(w, strconv.Itoa(i))
	} else {
		HttpError(w)
		return
	}
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	if shuttingDown {
		HttpOutput(w, "shutdown in progress...")
		return
	}

	if r.Method == http.MethodGet {
		stats := GetStats()
		json.NewEncoder(w).Encode(stats)
	}
}

func ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HttpError(w)
		return
	}

	shuttingDown = true
	LogMessage("attempting graceful shutdown")
	go func() {
		for {
			if !IsPending() {
				LogMessage("exiting")
				os.Exit(0)
			}
		}
	}()
	HttpOutput(w, "1")
}

func runServer() {
	http.HandleFunc("/hash", HashHandler)
	http.HandleFunc("/hash/", HashHandler)
	http.HandleFunc("/stats", StatsHandler)
	http.HandleFunc("/shutdown", ShutdownHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	LogMessage("starting server on port " + port)
	//make and wire up everything
	Store = make(map[int]Record)
	encryptJobs = make(chan IdAndValue)
	encryptResults = make(chan IdAndValue)
	go EncryptionWorker(encryptJobs, encryptResults)
	completeRecordJobs = make(chan IdAndValue)
	completeRecordResults = make(chan bool)
	go CompleteRecordWorker(completeRecordJobs, completeRecordResults)
	runServer()
}
