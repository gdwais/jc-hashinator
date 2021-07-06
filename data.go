package main

import "time"

//db
var Store map[int]Record

//channels
var encryptJobs chan IdAndValue
var encryptResults chan IdAndValue
var completeRecordJobs chan IdAndValue
var completeRecordResults chan bool

//misc
var port string = "8080"
var shuttingDown bool = false
var processCount int = 0

type Record struct {
	Hash           string
	Pending        bool
	Received       time.Time
	Completed      time.Time
	ProcessingTime time.Duration
}

type Stats struct {
	Total   int     `json:"total"`
	Average float64 `json:"average"`
}

type IdAndValue struct {
	Id    int
	Value string
}
