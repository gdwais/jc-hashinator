package main

import "time"

var Store map[int]Record

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

//public functions

func GetHash(i int) string {
	record, ok := Store[i]
	if ok {
		return record.Hash
	}
	return "no hash found"
}

func AddRecord() int {
	i := next()
	Store[i] = Record{Pending: true, Received: time.Now()}
	return i
}

func CompleteRecord(i int, hash string) {
	record, ok := Store[i]
	if ok {
		record.Hash = hash
		record.Completed = time.Now()
		record.ProcessingTime = record.Completed.Sub(record.Received)
		record.Pending = false
		Store[i] = record
	}
}

func GetStats() Stats {
	var length int = len(Store)
	durations := make([]time.Duration, 0, length)
	for _, v := range Store {
		durations = append(durations, v.ProcessingTime)
	}
	total := 0
	for _, duration := range durations {
		total = total + int(duration.Microseconds())
	}
	return Stats{Total: length, Average: float64(total) / float64(length)}
}

// private helpers

const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

func maximum() int {
	var maxNumber int
	if len(Store) == 0 {
		return maxNumber
	}
	maxNumber = minInt
	for n := range Store {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return maxNumber
}

func next() int {
	return maximum() + 1
}
