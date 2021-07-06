# jc-hashinator

Simple api (written in golang) for the purposes of hashing passwords 

## Startup

Startup the api with - api will run on port 8080
* go run .

You'll need to include the Authorization token "IW4nt2W0rk4JumpCl0ud" -  Otherwise all http requests will be rejected.  Reference the usage scripts for more info.

## Usage

I established some quick scripts for usage:

* **bin/load** : fills up the store with 200 records
* **bin/stats** : returns a json object with total number of records and average duration of processing
* **bin/shutdown** : triggers a graceful shutdown. all incoming other requests are rejected and if anything is in progress waits until complete
* **bin/hash x** : gets the hash value for the appropriate requested record - example: bin/hash 1

## Architecture Overview

### Overview

I wanted to give the application some basic structure without getting too bogged down in modules so I kept the application flat.  Separating code into manageable files helps to keep the application scaleable as functionality gets added.

#### **Main** 

Contains variable instatiation and wiring up the various routes with the various handlers as well as injecting the middleware.

#### **Data**

Contains all global variables and struct definitions.

#### **Store**

Contains all functions designed to insert, update or access the data store - if there was a database being leveraged here this would be the repository or query layer.

#### **Handlers**

Contains handler functions and middlware for processing requests.  This could do for some cleanup and refactoring currently.  Implementing a routing library would likely make it a bit more readable.

#### **Workers**

Contains worker functions that complete jobs - this pattern was helpful for me to understand channels and goroutines

#### **Logger**

This is preferential but I like to create centralized logging functions so that you can easily turn on/off logging or implement logging everything to a new service - ELK or slack or whatever you like.

## TODO 
* ~~establish endpoints~~
* ~~add store~~
* ~~add manager~~
* ~~wire up store to manager~~
* ~~wire up endpoints to manager/store~~
* ~~unified logging/output module~~
* ~~build out scripts for quick testing~~
* ~~implement channels~~
* ~~API Authentication via basic text key~~
* unit tests
* rework shutdown to use context  
* refactor main to use router 
* refactor into modules