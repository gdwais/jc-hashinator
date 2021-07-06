package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func HashHandler(w http.ResponseWriter, r *http.Request) {
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
			processCount = processCount + 1
			encryptJobs <- IdAndValue{Id: i, Value: password}
			encryptResults := <-encryptResults
			completeRecordJobs <- encryptResults
			complete := <-completeRecordResults
			if complete {
				LogMessage(strconv.Itoa(i) + " completed")
				processCount = processCount - 1
			}
		}()
		HttpOutput(w, strconv.Itoa(i))
	} else {
		HttpError(w)
		return
	}
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
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
		// this is hacky and i know there is a better way.  just ran out of time :)
		for {
			if processCount == 0 {
				LogMessage("exiting")
				os.Exit(0)
			}
		}
	}()
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if shuttingDown {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("awaiting shutdown..."))
		} else {
			authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
			if len(authHeader) != 2 {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Malformed Token"))
			} else {
				var token string = string(authHeader[1])

				if token == "IW4nt2W0rk4JumpCl0ud" {
					next.ServeHTTP(w, r)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Unauthorized"))
				}
			}
		}
	})
}
