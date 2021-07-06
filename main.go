package main

import (
	"log"
	"net/http"
)

func runServer() {
	http.Handle("/hash", Middleware(http.HandlerFunc(HashHandler)))
	http.Handle("/hash/", Middleware(http.HandlerFunc(HashHandler)))
	http.Handle("/stats", Middleware(http.HandlerFunc(StatsHandler)))
	http.Handle("/shutdown", Middleware(http.HandlerFunc(ShutdownHandler)))

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
	processCount = 0
	runServer()
}
