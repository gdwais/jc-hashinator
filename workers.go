package main

import (
	"crypto/sha512"
	"encoding/base64"
	"strconv"
	"time"
)

func EncryptionWorker(jobs <-chan IdAndValue, results chan<- IdAndValue) {
	for request := range jobs {
		LogMessage("starting work on encryption : " + strconv.Itoa(request.Id))
		go func(req IdAndValue) {
			time.Sleep(5 * time.Second)
			hasher := sha512.New()
			hasher.Write([]byte(req.Value))
			var hashedPasswordBytes = hasher.Sum(nil)
			hash := base64.URLEncoding.EncodeToString(hashedPasswordBytes)
			LogMessage("completed work on encryption : " + hash)
			results <- IdAndValue{Id: req.Id, Value: hash}
		}(request)
	}
}

func CompleteRecordWorker(jobs <-chan IdAndValue, results chan<- bool) {
	for request := range jobs {
		LogMessage("starting work on complete record : " + strconv.Itoa(request.Id))
		result := CompleteRecord(request)
		LogMessage("completed work on complete record : " + strconv.Itoa(request.Id))
		results <- result
	}
}
