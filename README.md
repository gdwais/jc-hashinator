# jc-hashinator

simple api (written in golang)

## Startup
startup the api with - api will run on port 8080
* go run .

## Usage

established some quick scripts for usage 

* **bin/load** : fills up the store with 200 records
* **bin/stats** : returns a json object with total number of records and average duration of processing
* **bin/shutdown** : triggers a graceful shutdown. all incoming other requests are rejected and if anything is in progress waits until complete
* **bin/hash x** : gets the hash value for the appropriate requested record - example: bin/hash 1

## Architecture Overview

* **Main** : endpoint and handler definitions
* **Logger** : unified logging functions
* **Store** : data array and data manipulation functions
* **Manager** : business logic

## TODO 
* ~~establish endpoints~~
* ~~add store~~
* ~~add manager~~
* ~~wire up store to manager~~
* ~~wire up endpoints to manager/store~~
* ~~unified logging/output module~~
* ~~build out scripts for quick testing~~
* implement channels
* refactor shutdown to use channel
* API Authentication via JWT or basic text key
* unit tests
* refactor main to use router 
* refactor into modules