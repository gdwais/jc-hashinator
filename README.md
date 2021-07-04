# jc-hashinator

simple api (written in golang)

## Endpoints

### GET - /stats
```
// return stat json
{
    "total": 1,  //int total of all passwords
    "average": 123 //average time of processing
}
```

### POST - /hash

payload example
```
password=passwordToHash
```
response example
```
// return int identifier
```

### GET - /hash/{id}
response example
```
//return hashed password value
```

### GET - shutdown

* new requests are rejected
* wait for pending work to complete and then exit



## Architecture Overview
* **Main** : endpoint and handler definitions
* **Logger** : unified logging functions
* **Auth** : authorization related functions
* **Store** : data array and data manipulation functions
* **Manager** : business logic

## TODO - may not get to
### Production quality `must haves
* API Authentication via JWT
* unified logging module
* 