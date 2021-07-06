package main

import (
	"time"
)

//public functions

func GetHash(i int) string {
	record, ok := Store[i]
	if ok {
		return record.Hash
	}
	return "no hash found"
}

func AddRecord() int {
	id := next()
	Store[id] = Record{Pending: true, Received: time.Now()}
	return id
}

func CompleteRecord(request IdAndValue) bool {
	record, ok := Store[request.Id]
	if ok {
		record.Hash = request.Value
		record.Completed = time.Now()
		record.ProcessingTime = record.Completed.Sub(record.Received)
		record.Pending = false
		Store[request.Id] = record
	}
	return true
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
	if total > 0 {
		return Stats{Total: length, Average: float64(total) / float64(length)}
	} else {
		return Stats{Total: 0, Average: 0}
	}
}

func IsPending() bool {
	for _, v := range Store {
		if v.Pending {
			return true
		}
	}
	return false
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
