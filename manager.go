package main

import (
	"crypto/sha512"
	"encoding/base64"
	"time"
)

func ProcessAsyncRequest(i int, password string) {
	go func() {
		time.Sleep(5 * time.Second)
		hasher := sha512.New()
		hasher.Write([]byte(password))
		var hashedPasswordBytes = hasher.Sum(nil)
		var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)
		CompleteRecord(i, base64EncodedPasswordHash)
	}()
}
